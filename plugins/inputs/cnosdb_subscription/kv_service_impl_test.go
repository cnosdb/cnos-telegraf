package cnosdb_subscription

import (
	"encoding/json"
	"testing"

	"github.com/influxdata/telegraf/plugins/inputs/cnosdb_subscription/cnosdb"
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
		`{ "column_type": { "Field": { "Geometry": { "sub_type": "Point", "srid": 10 } } } }`,
		`{ "column_type": { "Field": { "Geometry": { } } } }`,
	}
	var expectedColType = []cnosdb.ColumnTypeUnited{
		{ColumnType: cnosdb.ColumnTypeUnknown, TimeUnit: cnosdb.TimeUnitUnknown},
		{ColumnType: cnosdb.ColumnTypeTag, TimeUnit: cnosdb.TimeUnitUnknown},
		{ColumnType: cnosdb.ColumnTypeUnknown, TimeUnit: cnosdb.TimeUnitUnknown},
		{ColumnType: cnosdb.ColumnTypeTime, TimeUnit: cnosdb.TimeUnitUnknown},
		{ColumnType: cnosdb.ColumnTypeTime, TimeUnit: cnosdb.TimeUnitSecond},
		{ColumnType: cnosdb.ColumnTypeTime, TimeUnit: cnosdb.TimeUnitMillisecond},
		{ColumnType: cnosdb.ColumnTypeTime, TimeUnit: cnosdb.TimeUnitMicrosecond},
		{ColumnType: cnosdb.ColumnTypeTime, TimeUnit: cnosdb.TimeUnitNanosecond},
		{ColumnType: cnosdb.ColumnTypeUnknown, TimeUnit: cnosdb.TimeUnitUnknown},
		{ColumnType: cnosdb.ColumnTypeFieldUnknown, TimeUnit: cnosdb.TimeUnitUnknown},
		{ColumnType: cnosdb.ColumnTypeFieldFloat, TimeUnit: cnosdb.TimeUnitUnknown},
		{ColumnType: cnosdb.ColumnTypeFieldInteger, TimeUnit: cnosdb.TimeUnitUnknown},
		{ColumnType: cnosdb.ColumnTypeFieldUnsigned, TimeUnit: cnosdb.TimeUnitUnknown},
		{ColumnType: cnosdb.ColumnTypeFieldBoolean, TimeUnit: cnosdb.TimeUnitUnknown},
		{ColumnType: cnosdb.ColumnTypeFieldString, TimeUnit: cnosdb.TimeUnitUnknown},
		{ColumnType: cnosdb.ColumnTypeFieldGeometry, TimeUnit: cnosdb.TimeUnitUnknown},
		{ColumnType: cnosdb.ColumnTypeFieldGeometry, TimeUnit: cnosdb.TimeUnitUnknown},
	}

	if len(jsonStrings) != len(expectedColType) {
		t.Fatal("Incorrect init of test case")
	}

	for i := 0; i < len(jsonStrings); i++ {
		jsonStr := jsonStrings[i]
		obj := cnosdb.TableColumn{}
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
