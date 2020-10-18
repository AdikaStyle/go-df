package dataframe

import (
	"github.com/AdikaStyle/go-df/backend"
	"github.com/AdikaStyle/go-df/conds"
	"github.com/AdikaStyle/go-df/types"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestDefaultTransformable_UpdateColumn(t *testing.T) {
	df := newColumnarDataframe(transBackend())

	df.UpdateColumn("inventory", func(id int, row backend.Row) types.TypedValue {
		return 1000 * row["inventory"].(types.Integer)
	})

	df.VisitRows(func(id int, row backend.Row) {
		assert.True(t, row["inventory"].(types.Integer) > 10000)
	})

	df.Print(os.Stdout)
}

func TestDefaultTransformable_AddColumn(t *testing.T) {
	df := newColumnarDataframe(transBackend())

	df.AddColumn("isApple", types.KindBoolean, func(id int, row backend.Row) types.TypedValue {
		pname := row["name"].String()
		if pname == "Macbook" || pname == "iPad" {
			return types.Boolean(true)
		}
		return types.Boolean(false)
	})

	df.VisitRows(func(id int, row backend.Row) {
		pname := row["name"].String()
		if pname == "Macbook" || pname == "iPad" {
			assert.True(t, row["isApple"].Equals(types.Boolean(true)))
		} else {
			assert.True(t, row["isApple"].Equals(types.Boolean(false)))
		}
	})

	df.Print(os.Stdout)
}

func TestDefaultTransformable_Concat(t *testing.T) {
	df := newColumnarDataframe(transBackend())

	a0a1, a2a3 := df.Split(conds.In("SKU", types.String("A0"), types.String("A1")))

	a0a1.Concat(a2a3)

	a0a1.VisitRows(func(id int, row backend.Row) {
		assert.EqualValues(t, a0a1.getBackend().GetRow(id), df.getBackend().GetRow(id))
	})

	a0a1.Print(os.Stdout)
}

func TestDefaultTransformable_Split(t *testing.T) {
	df := newColumnarDataframe(transBackend())

	a0a1, a2a3 := df.Split(conds.In("SKU", types.String("A0"), types.String("A1")))

	assert.EqualValues(t, 3, len(a0a1.GetHeaders()))
	assert.EqualValues(t, 2, a0a1.GetRowCount())
	assert.EqualValues(t, "A0", a0a1.getBackend().GetRow(0)["SKU"])
	assert.EqualValues(t, "A1", a0a1.getBackend().GetRow(1)["SKU"])

	assert.EqualValues(t, 3, len(a2a3.GetHeaders()))
	assert.EqualValues(t, 2, a2a3.GetRowCount())
	assert.EqualValues(t, "A2", a2a3.getBackend().GetRow(0)["SKU"])
	assert.EqualValues(t, "A3", a2a3.getBackend().GetRow(1)["SKU"])
}

func transBackend() backend.Backend {
	headers := backend.Headers{
		{Name: "SKU", Kind: types.KindString},
		{Name: "name", Kind: types.KindString},
		{Name: "inventory", Kind: types.KindInteger},
	}
	be := backend.NewColumnarBackend(headers)

	be.AppendRow(
		backend.Row{"SKU": types.String("A0"), "name": types.String("Macbook"), "inventory": types.Integer(323)},
		backend.Row{"SKU": types.String("A1"), "name": types.String("PC"), "inventory": types.Integer(1000)},
		backend.Row{"SKU": types.String("A2"), "name": types.String("PSP"), "inventory": types.Integer(203)},
		backend.Row{"SKU": types.String("A3"), "name": types.String("iPad"), "inventory": types.Integer(120)},
	)
	return be
}
