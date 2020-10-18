package dataframe

import (
	"github.com/AdikaStyle/go-df/backend"
	"github.com/AdikaStyle/go-df/conds"
	"github.com/AdikaStyle/go-df/types"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestDefaultJoinable_InnerJoin(t *testing.T) {
	sales := newColumnarDataframe(aggBackend())
	catalog := newColumnarDataframe(transBackend())

	joined := sales.InnerJoin(catalog, conds.On("SKU"))

	assert.EqualValues(t, aggBackend().GetRowCount(), joined.GetRowCount())
	assert.EqualValues(t, aggBackend().GetRowCount(), joined.GetRowCount())
	assert.EqualValues(t, len(sales.GetHeaders())+len(catalog.GetHeaders())-1, len(joined.GetHeaders()))

	joined.VisitRows(func(id int, row backend.Row) {
		if row["SKU"] == types.String("A0") {
			l := sales.getBackend().GetRow(id)
			r := catalog.getBackend().GetRow(0)
			c := combineRows(l, r, false)
			assert.EqualValues(t, c, row)
		} else if row["SKU"] == types.String("A1") {
			l := sales.getBackend().GetRow(id)
			r := catalog.getBackend().GetRow(1)
			c := combineRows(l, r, false)
			assert.EqualValues(t, c, row)
		} else {
			assert.FailNow(t, "")
		}
	})

	joined.Print(os.Stdout)
}

func TestDefaultJoinable_LeftJoin(t *testing.T) {
	sales := newColumnarDataframe(aggBackend())
	catalog := newColumnarDataframe(transBackend()).Filter(conds.Eq("SKU", types.String("A1")))

	joined := sales.LeftJoin(catalog, conds.On("SKU"))

	assert.EqualValues(t, aggBackend().GetRowCount(), joined.GetRowCount())
	assert.EqualValues(t, aggBackend().GetRowCount(), joined.GetRowCount())
	assert.EqualValues(t, len(sales.GetHeaders())+len(catalog.GetHeaders())-1, len(joined.GetHeaders()))

	joined.VisitRows(func(id int, row backend.Row) {
		if row["SKU"] == types.String("A0") {
			l := sales.getBackend().GetRow(id)
			assert.EqualValues(t, types.Missing, row["inventory"])
			assert.EqualValues(t, types.Missing, row["name"])
			delete(row, "inventory")
			delete(row, "name")
			assert.EqualValues(t, l, row)
		} else if row["SKU"] == types.String("A1") {
			l := sales.getBackend().GetRow(id)
			r := catalog.getBackend().GetRow(1)
			c := combineRows(l, r, false)
			assert.EqualValues(t, c, row)
		} else {
			assert.FailNow(t, "")
		}
	})

	joined.Print(os.Stdout)
}

func TestDefaultJoinable_OuterJoin(t *testing.T) {

}
