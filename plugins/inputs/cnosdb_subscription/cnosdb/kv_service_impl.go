package cnosdb

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	flatbuffers "github.com/google/flatbuffers/go"
	"google.golang.org/grpc"
	"io"
	"math"
	"time"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
	"github.com/apache/arrow/go/v15/arrow/ipc"
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/metric"
	"github.com/influxdata/telegraf/plugins/inputs/cnosdb_subscription/cnosdb/generated/kv_service"
	models_v3 "github.com/influxdata/telegraf/plugins/inputs/cnosdb_subscription/cnosdb/generated/models/v3"
	cnosdb_v3 "github.com/influxdata/telegraf/plugins/inputs/cnosdb_subscription/cnosdb/v3"
	cnosdb_v4 "github.com/influxdata/telegraf/plugins/inputs/cnosdb_subscription/cnosdb/v4"
)

func RegisterTSKVServiceServer(s grpc.ServiceRegistrar, srv kv_service.TSKVServiceServer) {
	kv_service.RegisterTSKVServiceServer(s, srv)
}

var _ kv_service.TSKVServiceServer = (*TSKVServiceServerImpl)(nil)

type TSKVServiceServerImpl struct {
	accumulator telegraf.Accumulator

	kv_service.UnimplementedTSKVServiceServer
}

func NewTSKVService(acc telegraf.Accumulator) kv_service.TSKVServiceServer {
	return TSKVServiceServerImpl{
		accumulator: acc,
	}
}

// WritePoints receive subscription requests with flat-buffers message `Points`,
// parse them into metrics and write to accumulator.
func (s TSKVServiceServerImpl) WritePoints(server kv_service.TSKVService_WritePointsServer) error {
	for {
		req, err := server.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			s.accumulator.AddError(fmt.Errorf("failed to receive WritePointsRequest: %w", err))
			return server.Send(&kv_service.WritePointsResponse{
				PointsNumber: 0,
			})
		}

		var points models_v3.Points
		flatbuffers.GetRootAs(req.Points, 0, &points)

		var fbTable models_v3.Table
		var fbSchema models_v3.Schema
		var fbPoint models_v3.Point
		var fbTag models_v3.Tag
		var fbField models_v3.Field
		for tableIndex := 0; tableIndex < points.TablesLength(); tableIndex++ {
			if !points.Tables(&fbTable, tableIndex) {
				continue
			}
			fbTable.Schema(&fbSchema)
			table := string(fbTable.TabBytes())
			for pointIndex := 0; pointIndex < fbTable.PointsLength(); pointIndex++ {
				if !fbTable.Points(&fbPoint, pointIndex) {
					continue
				}

				tags := make(map[string]string)
				fields := make(map[string]interface{})

				tagsBitSet := cnosdb_v3.BitSet{
					Buf: fbPoint.TagsNullbitBytes(),
					Len: fbSchema.TagNameLength(),
				}
				if fbPoint.TagsLength() > 0 {
					for tagIndex := 0; tagIndex < fbPoint.TagsLength(); tagIndex++ {
						if !tagsBitSet.Get(tagIndex) {
							continue
						}
						if !fbPoint.Tags(&fbTag, tagIndex) {
							continue
						}
						tags[string(fbSchema.TagName(tagIndex))] = string(fbTag.ValueBytes())
					}
				}
				fieldsBitSet := cnosdb_v3.BitSet{
					Buf: fbPoint.FieldsNullbitBytes(),
					Len: fbSchema.FieldNameLength(),
				}
				if fbPoint.FieldsLength() > 0 {
					for fieldIndex := 0; fieldIndex < fbPoint.FieldsLength(); fieldIndex++ {
						if !fieldsBitSet.Get(fieldIndex) {
							continue
						}
						if !fbPoint.Fields(&fbField, fieldIndex) {
							continue
						}
						switch fbSchema.FieldType(fieldIndex) {
						case models_v3.FieldTypeInteger:
							v := binary.BigEndian.Uint64(fbField.ValueBytes())
							fields[string(fbSchema.FieldName(fieldIndex))] = int64(v)
						case models_v3.FieldTypeUnsigned:
							v := binary.BigEndian.Uint64(fbField.ValueBytes())
							fields[string(fbSchema.FieldName(fieldIndex))] = v
						case models_v3.FieldTypeFloat:
							tmp := binary.BigEndian.Uint64(fbField.ValueBytes())
							v := math.Float64frombits(tmp)
							fields[string(fbSchema.FieldName(fieldIndex))] = v
						case models_v3.FieldTypeBoolean:
							v := fbField.ValueBytes()
							fields[string(fbSchema.FieldName(fieldIndex))] = v[0] == byte(1)
						case models_v3.FieldTypeString:
							v := string(fbField.ValueBytes())
							fields[string(fbSchema.FieldName(fieldIndex))] = v
						default:
							// Do nothing ?
						}
					}
				}
				s.accumulator.AddFields(table, fields, tags)
			}

		}

	}

	return nil
}

// WriteSubscription receive subscription request that has a CnosDB table schema and an Arrow RecordBatch,
// parse them into metrics and write to accumulator.
func (s TSKVServiceServerImpl) WriteSubscription(server kv_service.TSKVService_WriteSubscriptionServer) error {
	for {
		req, err := server.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			s.accumulator.AddError(fmt.Errorf("failed to receive SubscriptionRequest: %w", err))
			return server.Send(&kv_service.SubscriptionResponse{})
		}

		var tableSchema cnosdb_v4.TskvTableSchema
		if err := json.Unmarshal(req.TableSchema, &tableSchema); err != nil {
			return server.Send(&kv_service.SubscriptionResponse{})
		}

		recordReader, err := ipc.NewReader(bufio.NewReader(bytes.NewReader(req.RecordData)))
		if err != nil {
			s.accumulator.AddError(fmt.Errorf("failed to read record data: %w", err))
			return server.Send(&kv_service.SubscriptionResponse{})
		}

		columnIndexToType := make([]cnosdb_v4.ColumnTypeUnited, len(tableSchema.Columns))
		columnIndexToName := make([]string, len(tableSchema.Columns))

		for _, col := range tableSchema.Columns {
			colType := col.GetColumnTypeUnited()
			if colType.ColumnType != cnosdb_v4.ColumnTypeUnknown && colType.ColumnType != cnosdb_v4.ColumnTypeFieldUnknown {
				columnIndexToType[col.ID] = colType
			} else {
				s.accumulator.AddError(fmt.Errorf("column '%s': type is unknown: ", col.Name))
				return server.Send(&kv_service.SubscriptionResponse{})
			}
			columnIndexToName[col.ID] = col.Name
		}
		for {
			r, err := recordReader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				s.accumulator.AddError(fmt.Errorf("failed to read record data: %w", err))
				continue
			}

			numRows := r.NumRows()
			numCols := r.NumCols()
			colArrays := make([]arrow.Array, len(tableSchema.Columns))

			for j, col := range r.Columns() {
				colType := columnIndexToType[j]
				switch colType.ColumnType {
				case cnosdb_v4.ColumnTypeTag:
					colArrays[j] = array.NewStringData(col.Data())
				case cnosdb_v4.ColumnTypeTime:
					if colType.TimeUnit == cnosdb_v4.TimeUnitUnknown {
						colName := columnIndexToName[j]
						s.accumulator.AddError(fmt.Errorf("column '%s': parsed time unit(%d) is unknown", colName, colType))
						return server.Send(&kv_service.SubscriptionResponse{})
					}
					colArrays[j] = array.NewTime64Data(col.Data())
				case cnosdb_v4.ColumnTypeFieldFloat:
					colArrays[j] = array.NewFloat64Data(col.Data())
				case cnosdb_v4.ColumnTypeFieldInteger:
					colArrays[j] = array.NewInt64Data(col.Data())
				case cnosdb_v4.ColumnTypeFieldUnsigned:
					colArrays[j] = array.NewUint64Data(col.Data())
				case cnosdb_v4.ColumnTypeFieldBoolean:
					colArrays[j] = array.NewBooleanData(col.Data())
				case cnosdb_v4.ColumnTypeFieldString, cnosdb_v4.ColumnTypeFieldGeometry:
					colArrays[j] = array.NewStringData(col.Data())
				default:
					colName := columnIndexToName[j]
					s.accumulator.AddError(fmt.Errorf("column '%s': parsed type(%d) is unknown", colName, colType))
					return server.Send(&kv_service.SubscriptionResponse{})
				}
			}

			for i := 0; i < int(numRows); i++ {
				tags := make(map[string]string)
				fields := make(map[string]interface{})
				var timestamp time.Time
				hasTimestamp := false
				for j := 0; j < int(numCols); j++ {
					if colArrays[j].IsNull(i) {
						continue
					}
					colType := columnIndexToType[j]
					switch colType.ColumnType {
					case cnosdb_v4.ColumnTypeTag:
						tags[columnIndexToName[j]] = colArrays[j].(*array.String).Value(i)
					case cnosdb_v4.ColumnTypeTime:
						switch colType.TimeUnit {
						case cnosdb_v4.TimeUnitSecond:
							hasTimestamp = true
							timestamp = time.Unix(int64(colArrays[j].(*array.Time64).Value(i)), 0)
						case cnosdb_v4.TimeUnitMillisecond:
							hasTimestamp = true
							timestamp = time.UnixMilli(int64(colArrays[j].(*array.Time64).Value(i)))
						case cnosdb_v4.TimeUnitMicrosecond:
							hasTimestamp = true
							timestamp = time.UnixMicro(int64(colArrays[j].(*array.Time64).Value(i)))
						case cnosdb_v4.TimeUnitNanosecond:
							hasTimestamp = true
							timestamp = time.Unix(0, int64(colArrays[j].(*array.Time64).Value(i)))
						default:
							colName := columnIndexToName[j]
							s.accumulator.AddError(fmt.Errorf("column '%s': parsed time unit(%d) is unknown", colName, colType))
							return server.Send(&kv_service.SubscriptionResponse{})
						}
					case cnosdb_v4.ColumnTypeFieldFloat:
						fields[columnIndexToName[j]] = colArrays[j].(*array.Float64).Value(i)
					case cnosdb_v4.ColumnTypeFieldInteger:
						fields[columnIndexToName[j]] = colArrays[j].(*array.Int64).Value(i)
					case cnosdb_v4.ColumnTypeFieldUnsigned:
						fields[columnIndexToName[j]] = colArrays[j].(*array.Uint64).Value(i)
					case cnosdb_v4.ColumnTypeFieldBoolean:
						fields[columnIndexToName[j]] = colArrays[j].(*array.Boolean).Value(i)
					case cnosdb_v4.ColumnTypeFieldString, cnosdb_v4.ColumnTypeFieldGeometry:
						fields[columnIndexToName[j]] = colArrays[j].(*array.String).Value(i)
					default:
						colName := columnIndexToName[j]
						s.accumulator.AddError(fmt.Errorf("column '%s': parsed type(%d) is unknown", colName, colType))
						return server.Send(&kv_service.SubscriptionResponse{})
					}
				}
				switch acc := s.accumulator.(type) {
				case telegraf.HighPriorityAccumulator:
					if hasTimestamp {
						if err = acc.AddMetricHighPriority(metric.New(tableSchema.Name, tags, fields, timestamp)); err != nil {
							acc.AddError(fmt.Errorf("writing data to output failed: %w", err))
						}
					} else {
						if err = acc.AddMetricHighPriority(metric.New(tableSchema.Name, tags, fields, time.Now())); err != nil {
							acc.AddError(fmt.Errorf("writing data to output failed: %w", err))
						}
					}
				default:
					if hasTimestamp {
						acc.AddMetric(metric.New(tableSchema.Name, tags, fields, timestamp))
					} else {
						acc.AddFields(tableSchema.Name, fields, tags)
					}
				}
			}
		}
	}
	return nil
}
