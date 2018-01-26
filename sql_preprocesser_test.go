package sqlutil

import (
	"testing"
)

func TestRemoveComments(t *testing.T) {
	for _, sql := range testSQLs {
		// t.Logf("sql:%v", sql)
		n, err := RemoveComments(sql)
		if err != nil {
			t.Fatalf(err.Error())
		}
		t.Logf("sql:%v, new sql:%v", sql, n)
	}
}
