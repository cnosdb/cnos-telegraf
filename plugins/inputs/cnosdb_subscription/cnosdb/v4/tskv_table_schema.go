package v4

type ColumnTypeCode uint32

const (
	ColumnTypeUnknown ColumnTypeCode = iota
	ColumnTypeTag
	ColumnTypeTime
	ColumnTypeFieldUnknown
	ColumnTypeFieldFloat
	ColumnTypeFieldInteger
	ColumnTypeFieldUnsigned
	ColumnTypeFieldBoolean
	ColumnTypeFieldString
	ColumnTypeFieldGeometry
)

type TimeUnitCode uint32

const (
	TimeUnitUnknown TimeUnitCode = iota
	TimeUnitSecond
	TimeUnitMillisecond
	TimeUnitMicrosecond
	TimeUnitNanosecond
)

type TskvTableSchema struct {
	// Tenant string `json:"tenant"`
	// Db   string `json:"db"`
	Name string `json:"name"`
	// SchemaId uint64 `json:"schema_id"` // v2.4.0 field
	// SchemaVersion uint64 `json:"schema_version"` // v2.4.1 field
	// NextColumnID uint32 `json:"next_column_id"`
	Columns []TableColumn `json:"columns"`
	// ColumnsIndex  map[string]uint32 `json:"columns_index"`
}

type TableColumn struct {
	ID         uint64      `json:"id"`
	Name       string      `json:"name"`
	ColumnType interface{} `json:"column_type"`
	// Encoding interface{} `json:"encoding"`
}

type ColumnTypeUnited struct {
	ColumnType ColumnTypeCode
	TimeUnit   TimeUnitCode
}

func (c *TableColumn) GetColumnTypeUnited() ColumnTypeUnited {
	switch columnType := c.ColumnType.(type) {
	case string:
		if columnType == "Tag" {
			// "column_type": "Tag"
			return ColumnTypeUnited{
				ColumnType: ColumnTypeTag,
				TimeUnit:   TimeUnitUnknown,
			}
		} else {
			// In cnosdb-v2.4.0, columnType is string
			// After cnosdb-v2.4.0, columnType is string or object
			columnTypeCode := ColumnTypeUnknown
			timeUnitCode := TimeUnitUnknown
			switch columnType {
			case "TAG_STRING":
				// "column_type": "TAG_STRING"
				return ColumnTypeUnited{
					ColumnType: ColumnTypeTag,
					TimeUnit:   TimeUnitUnknown,
				}
			case "FIELD_STRING":
				// "column_type": "FIELD_STRING"
				columnTypeCode = ColumnTypeFieldString
			case "FIELD_BIGINT":
				// "column_type": "FIELD_BIGINT""
				columnTypeCode = ColumnTypeFieldInteger
			case "FIELD_BIGINT UNSIGNED":
				// "column_type": "FIELD_BIGINT UNSIGNED""
				columnTypeCode = ColumnTypeFieldUnsigned
			case "FIELD_DOUBLE":
				// "column_type": "FIELD_STRING"
				columnTypeCode = ColumnTypeFieldFloat
			case "FIELD_BOOLEAN":
				// "column_type": "FIELD_BOOLEAN""
				columnTypeCode = ColumnTypeFieldBoolean
			case "TIME_TIMESTAMP(SECOND)":
				// "column_type": "TIME_TIMESTAMP(SECOND)""
				columnTypeCode = ColumnTypeTime
				timeUnitCode = TimeUnitSecond
			case "TIME_TIMESTAMP(MILLISECOND)":
				// "column_type": "TIME_TIMESTAMP(MILLISECOND)""
				columnTypeCode = ColumnTypeTime
				timeUnitCode = TimeUnitMillisecond
			case "TIME_TIMESTAMP(MICROSECOND)":
				// "column_type": "TIME_TIMESTAMP(MICROSECOND)""
				columnTypeCode = ColumnTypeTime
				timeUnitCode = TimeUnitMicrosecond
			case "TIME_TIMESTAMP(NANOSECOND)":
				// "column_type": "TIME_TIMESTAMP(NANOSECOND)""
				columnTypeCode = ColumnTypeTime
				timeUnitCode = TimeUnitNanosecond
			}
			return ColumnTypeUnited{
				ColumnType: columnTypeCode,
				TimeUnit:   timeUnitCode,
			}
		}
	case map[string]interface{}:
		if timeUnitObj := columnType["Time"]; timeUnitObj != nil {
			// "column_type": {"Time":"Microsecond"}
			if timeUnit, ok := timeUnitObj.(string); ok {
				timeUnitCode := TimeUnitUnknown
				switch timeUnit {
				case "Second":
					timeUnitCode = TimeUnitSecond
				case "Millisecond":
					timeUnitCode = TimeUnitMillisecond
				case "Microsecond":
					timeUnitCode = TimeUnitMicrosecond
				case "Nanosecond":
					timeUnitCode = TimeUnitNanosecond
				}
				return ColumnTypeUnited{
					ColumnType: ColumnTypeTime,
					TimeUnit:   timeUnitCode,
				}
			}
		} else if fieldTypeObj := columnType["Field"]; fieldTypeObj != nil {
			fieldTypeCode := ColumnTypeFieldUnknown
			switch fieldType := fieldTypeObj.(type) {
			case string:
				switch fieldType {
				case "Float":
					// "column_type": {"Field":"Float"}
					fieldTypeCode = ColumnTypeFieldFloat
				case "Integer":
					// "column_type": {"Field":"Integer"}
					fieldTypeCode = ColumnTypeFieldInteger
				case "Unsigned":
					// "column_type": {"Field":"Unsigned"}
					fieldTypeCode = ColumnTypeFieldUnsigned
				case "Boolean":
					// "column_type": {"Field":"Boolean"}
					fieldTypeCode = ColumnTypeFieldBoolean
				case "String":
					// "column_type": {"Field":"String"}
					fieldTypeCode = ColumnTypeFieldString
				case "Geometry":
					// "column_type": {"Field":"Geometry"}
					fieldTypeCode = ColumnTypeFieldGeometry
				case "Unknown":
					// "column_type": {"Field":"Unknown"}
					fieldTypeCode = ColumnTypeFieldUnknown
				}
			case map[string]interface{}:
				if geometryInfo := fieldType["Geometry"]; geometryInfo != nil {
					// "column_type": {"Field":{"Geometry":{"sub_type":"Point"}}}
					fieldTypeCode = ColumnTypeFieldGeometry
				}
			}
			return ColumnTypeUnited{
				ColumnType: fieldTypeCode,
				TimeUnit:   TimeUnitUnknown,
			}
		}
	}

	return ColumnTypeUnited{
		ColumnType: ColumnTypeUnknown,
		TimeUnit:   TimeUnitUnknown,
	}
}
