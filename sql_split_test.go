package sqlutil

import "testing"

func TestSplitSQL(t *testing.T) {
	for _, sql := range testSQLs {
		t.Logf("sql:%v\n", sql)
		ss, err := Split(sql)
		if err != nil {
			t.Fatalf(err.Error())
		}

		for idx, s := range ss {
			t.Logf("%dth sql:%#v\n", idx, s)
		}
	}
}
