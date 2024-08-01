package cnosdb_subscription

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
	"github.com/apache/arrow/go/v15/arrow/ipc"
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/metric"
	"github.com/influxdata/telegraf/plugins/inputs/cnosdb_subscription/cnosdb"
	"github.com/influxdata/telegraf/plugins/inputs/cnosdb_subscription/protos"
)

//go:generate flatc -o ./protos --go --go-namespace protos --gen-onefile ./protos/flatbuffers/models.fbs
//go:generate protoc --go_out=./protos --go-grpc_out=./protos ./protos/protobuf/kv_service.proto

var _ protos.TSKVServiceServer = (*TSKVServiceServerImpl)(nil)

type TSKVServiceServerImpl struct {
	accumulator telegraf.Accumulator

	protos.UnimplementedTSKVServiceServer
}

func NewTSKVService(acc telegraf.Accumulator) protos.TSKVServiceServer {
	return TSKVServiceServerImpl{
		accumulator: acc,
	}
}

// WriteSubscription receive subscription request that has a CnosDB table schema and an Arrow RecordBatch,
// parse them into metrics and write to accumulator.
func (s TSKVServiceServerImpl) WriteSubscription(server protos.TSKVService_WriteSubscriptionServer) error {
	for {
		req, err := server.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			s.accumulator.AddError(fmt.Errorf("failed to receive SubscriptionRequest: %w", err))
			return server.Send(&protos.SubscriptionResponse{})
		}

		var tableSchema cnosdb.TskvTableSchema
		if err := json.Unmarshal(req.TableSchema, &tableSchema); err != nil {
			return server.Send(&protos.SubscriptionResponse{})
		}

		recordReader, err := ipc.NewReader(bufio.NewReader(bytes.NewReader(req.RecordData)))
		if err != nil {
			s.accumulator.AddError(fmt.Errorf("failed to read record data: %w", err))
			return server.Send(&protos.SubscriptionResponse{})
		}

		columnIndexToType := make([]cnosdb.ColumnTypeUnited, len(tableSchema.Columns))
		columnIndexToName := make([]string, len(tableSchema.Columns))

		for _, col := range tableSchema.Columns {
			colType := col.GetColumnTypeUnited()
			if colType.ColumnType != cnosdb.ColumnTypeUnknown && colType.ColumnType != cnosdb.ColumnTypeFieldUnknown {
				columnIndexToType[col.ID] = colType
			} else {
				s.accumulator.AddError(fmt.Errorf("column '%s': type is unknown: ", col.Name))
				return server.Send(&protos.SubscriptionResponse{})
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
				case cnosdb.ColumnTypeTag:
					colArrays[j] = array.NewStringData(col.Data())
				case cnosdb.ColumnTypeTime:
					if colType.TimeUnit == cnosdb.TimeUnitUnknown {
						colName := columnIndexToName[j]
						s.accumulator.AddError(fmt.Errorf("column '%s': parsed time unit(%d) is unknown", colName, colType))
						return server.Send(&protos.SubscriptionResponse{})
					}
					colArrays[j] = array.NewTime64Data(col.Data())
				case cnosdb.ColumnTypeFieldFloat:
					colArrays[j] = array.NewFloat64Data(col.Data())
				case cnosdb.ColumnTypeFieldInteger:
					colArrays[j] = array.NewInt64Data(col.Data())
				case cnosdb.ColumnTypeFieldUnsigned:
					colArrays[j] = array.NewUint64Data(col.Data())
				case cnosdb.ColumnTypeFieldBoolean:
					colArrays[j] = array.NewBooleanData(col.Data())
				case cnosdb.ColumnTypeFieldString, cnosdb.ColumnTypeFieldGeometry:
					colArrays[j] = array.NewStringData(col.Data())
				default:
					colName := columnIndexToName[j]
					s.accumulator.AddError(fmt.Errorf("column '%s': parsed type(%d) is unknown", colName, colType))
					return server.Send(&protos.SubscriptionResponse{})
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
					case cnosdb.ColumnTypeTag:
						tags[columnIndexToName[j]] = colArrays[j].(*array.String).Value(i)
					case cnosdb.ColumnTypeTime:
						switch colType.TimeUnit {
						case cnosdb.TimeUnitSecond:
							hasTimestamp = true
							timestamp = time.Unix(int64(colArrays[j].(*array.Time64).Value(i)), 0)
						case cnosdb.TimeUnitMillisecond:
							hasTimestamp = true
							timestamp = time.UnixMilli(int64(colArrays[j].(*array.Time64).Value(i)))
						case cnosdb.TimeUnitMicrosecond:
							hasTimestamp = true
							timestamp = time.UnixMicro(int64(colArrays[j].(*array.Time64).Value(i)))
						case cnosdb.TimeUnitNanosecond:
							hasTimestamp = true
							timestamp = time.Unix(0, int64(colArrays[j].(*array.Time64).Value(i)))
						default:
							colName := columnIndexToName[j]
							s.accumulator.AddError(fmt.Errorf("column '%s': parsed time unit(%d) is unknown", colName, colType))
							return server.Send(&protos.SubscriptionResponse{})
						}
					case cnosdb.ColumnTypeFieldFloat:
						fields[columnIndexToName[j]] = colArrays[j].(*array.Float64).Value(i)
					case cnosdb.ColumnTypeFieldInteger:
						fields[columnIndexToName[j]] = colArrays[j].(*array.Int64).Value(i)
					case cnosdb.ColumnTypeFieldUnsigned:
						fields[columnIndexToName[j]] = colArrays[j].(*array.Uint64).Value(i)
					case cnosdb.ColumnTypeFieldBoolean:
						fields[columnIndexToName[j]] = colArrays[j].(*array.Boolean).Value(i)
					case cnosdb.ColumnTypeFieldString, cnosdb.ColumnTypeFieldGeometry:
						fields[columnIndexToName[j]] = colArrays[j].(*array.String).Value(i)
					default:
						colName := columnIndexToName[j]
						s.accumulator.AddError(fmt.Errorf("column '%s': parsed type(%d) is unknown", colName, colType))
						return server.Send(&protos.SubscriptionResponse{})
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
