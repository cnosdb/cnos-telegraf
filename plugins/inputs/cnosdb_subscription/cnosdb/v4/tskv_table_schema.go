package v4

type ColumnType uint32

const (
	ColumnTypeUnknown ColumnType = iota
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

type TimeUnit uint32

const (
	TimeUnitUnknown TimeUnit = iota
	TimeUnitSecond
	TimeUnitMillisecond
	TimeUnitMicrosecond
	TimeUnitNanosecond
)

type TskvTableSchema struct {
	Tenant        string            `json:"tenant"`
	Db            string            `json:"db"`
	Name          string            `json:"name"`
	SchemaVersion uint64            `json:"schema_version"`
	NextColumnID  uint32            `json:"next_column_id"`
	Columns       []TableColumn     `json:"columns"`
	ColumnsIndex  map[string]uint32 `json:"columns_index"`
}

type TableColumn struct {
	ID         uint64      `json:"id"`
	Name       string      `json:"name"`
	ColumnType interface{} `json:"column_type"`
	Encoding   interface{} `json:"encoding"`
}

type ColumnTypeUnited struct {
	ColumnType ColumnType
	TimeUnit   TimeUnit
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
