package v4

import (
	"encoding/json"
	"testing"
)

func TestParseColumnTypeUnited(t *testing.T) {
	var jsonStrings = []string{
		`{ "column_type": "tag" }`,
		`{ "column_type": "Tag" }`,
		`{ "column_type": { "time": "Second" } }`,
		`{ "column_type": { "Time": "sec" } }`,
		`{ "column_type": { "Time": "Second" } }`,
		`{ "column_type": { "Time": "Millisecond" } }`,
		`{ "column_type": { "Time": "Microsecond" } }`,
		`{ "column_type": { "Time": "Nanosecond" } }`,
		`{ "column_type": { "field": "Float" } }`,
		`{ "column_type": { "Field": "f64" } }`,
		`{ "column_type": { "Field": "Float" } }`,
		`{ "column_type": { "Field": "Integer" } }`,
		`{ "column_type": { "Field": "Unsigned" } }`,
		`{ "column_type": { "Field": "Boolean" } }`,
		`{ "column_type": { "Field": "String" } }`,
		`{ "column_type": { "Field": { "Geometry": { "sub_type": "Point" } } } }`,
		`{ "column_type": { "Field": { "Geometry": { } } } }`,
	}
	var expectedColType = []ColumnTypeUnited{
		{ColumnType: ColumnTypeUnknown, TimeUnit: TimeUnitUnknown},
		{ColumnType: ColumnTypeTag, TimeUnit: TimeUnitUnknown},
		{ColumnType: ColumnTypeUnknown, TimeUnit: TimeUnitUnknown},
		{ColumnType: ColumnTypeTime, TimeUnit: TimeUnitUnknown},
		{ColumnType: ColumnTypeTime, TimeUnit: TimeUnitSecond},
		{ColumnType: ColumnTypeTime, TimeUnit: TimeUnitMillisecond},
		{ColumnType: ColumnTypeTime, TimeUnit: TimeUnitMicrosecond},
		{ColumnType: ColumnTypeTime, TimeUnit: TimeUnitNanosecond},
		{ColumnType: ColumnTypeUnknown, TimeUnit: TimeUnitUnknown},
		{ColumnType: ColumnTypeFieldUnknown, TimeUnit: TimeUnitUnknown},
		{ColumnType: ColumnTypeFieldFloat, TimeUnit: TimeUnitUnknown},
		{ColumnType: ColumnTypeFieldInteger, TimeUnit: TimeUnitUnknown},
		{ColumnType: ColumnTypeFieldUnsigned, TimeUnit: TimeUnitUnknown},
		{ColumnType: ColumnTypeFieldBoolean, TimeUnit: TimeUnitUnknown},
		{ColumnType: ColumnTypeFieldString, TimeUnit: TimeUnitUnknown},
		{ColumnType: ColumnTypeFieldGeometry, TimeUnit: TimeUnitUnknown},
		{ColumnType: ColumnTypeFieldGeometry, TimeUnit: TimeUnitUnknown},
	}

	if len(jsonStrings) != len(expectedColType) {
		t.Fatal("Incorrect init of test case")
	}

	for i := 0; i < len(jsonStrings); i++ {
		jsonStr := jsonStrings[i]
		obj := TableColumn{}
		if err := json.Unmarshal([]byte(jsonStr), &obj); err != nil {
			t.Error(err)
		}
		colType := obj.GetColumnTypeUnited()
		expColType := expectedColType[i]
		if obj.GetColumnTypeUnited() != expColType {
			t.Fatal("expected", expColType, "got", colType)
		}
	}
}
