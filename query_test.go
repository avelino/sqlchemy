package sqlchemy

import (
	"testing"
)

func TestQuery(t *testing.T) {
	t.Run("query all fields", func(t *testing.T) {
		testReset()
		q := testTable.Query()
		want := "SELECT t1.col0, t1.col1 FROM `test` AS `t1`"
		testGotWant(t, q.String(), want)
	})

	t.Run("query selected fields", func(t *testing.T) {
		testReset()
		q := testTable.Query(testTable.Field("col0")).Equals("col1", 100)
		want := "SELECT t1.col0 FROM `test` AS `t1` WHERE `t1`.`col1` = ( ? )"
		testGotWant(t, q.String(), want)
	})

	t.Run("query selected fields from subquery", func(t *testing.T) {
		testReset()
		q := testTable.Query().SubQuery().Query(testTable.Field("col0")).Equals("col1", 100)
		want := "SELECT t1.col0 FROM (SELECT t1.col1 FROM `test` AS `t1`) AS `t2` WHERE `t2`.`col1` = ( ? )"
		testGotWant(t, q.String(), want)
	})
}

func TestCountQuery(t *testing.T) {
	testReset()

	q := testTable.Query()
	q.GroupBy("col0")
	q.Limit(8)
	q.Offset(10)
	cq := q.countQuery()
	want := "SELECT COUNT(*) AS `count` FROM (SELECT t1.col0, t1.col1 FROM `test` AS `t1` GROUP BY `t1`.`col0`) AS `t2`"
	testGotWant(t, cq.String(), want)
}
