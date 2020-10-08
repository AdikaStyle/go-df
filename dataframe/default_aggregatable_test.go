package dataframe

import (
	"github.com/AdikaStyle/go-df/aggs"
	"github.com/AdikaStyle/go-df/backend"
	"github.com/AdikaStyle/go-df/types"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestDefaultAggregatable_GroupBy_OrderBy(t *testing.T) {
	df := newColumnarDataframe(aggBackend())
	gdf := df.Group(
		By{"SKU", "zone"},
		Aggs{"quantity": aggs.Sum(), "income": aggs.Sum()},
	).OrderBy("income", backend.Desc)

	assert.EqualValues(t, gdf.GetRowCount(), 5)

	gdf.Print(os.Stdout)
}

func aggBackend() backend.Backend {
	headers := backend.Headers{
		{Name: "SKU", Kind: types.KindString},
		{Name: "zone", Kind: types.KindString},
		{Name: "quantity", Kind: types.KindInteger},
		{Name: "income", Kind: types.KindDecimal},
	}
	be := backend.NewColumnarBackend(headers)
	be.AppendRow(
		backend.Row{"SKU": types.String("A0"), "zone": types.String("IL"), "quantity": types.Integer(2), "income": types.Decimal(200)},
		backend.Row{"SKU": types.String("A0"), "zone": types.String("IL"), "quantity": types.Integer(1), "income": types.Decimal(100)},
		backend.Row{"SKU": types.String("A0"), "zone": types.String("US"), "quantity": types.Integer(1), "income": types.Decimal(100)},
		backend.Row{"SKU": types.String("A1"), "zone": types.String("IL"), "quantity": types.Integer(3), "income": types.Decimal(600)},
		backend.Row{"SKU": types.String("A1"), "zone": types.String("IL"), "quantity": types.Integer(4), "income": types.Decimal(800)},
		backend.Row{"SKU": types.String("A1"), "zone": types.String("EU"), "quantity": types.Integer(20), "income": types.Decimal(4000)},
		backend.Row{"SKU": types.String("A1"), "zone": types.String("EU"), "quantity": types.Integer(1), "income": types.Decimal(200)},
		backend.Row{"SKU": types.String("A1"), "zone": types.String("EU"), "quantity": types.Integer(3), "income": types.Decimal(600)},
		backend.Row{"SKU": types.String("A1"), "zone": types.String("US"), "quantity": types.Integer(1), "income": types.Decimal(200)},
		backend.Row{"SKU": types.String("A1"), "zone": types.String("US"), "quantity": types.Integer(1), "income": types.Decimal(200)},
		backend.Row{"SKU": types.String("A1"), "zone": types.String("US"), "quantity": types.Integer(322), "income": types.Decimal(322 * 200)},
		backend.Row{"SKU": types.String("A1"), "zone": types.String("US"), "quantity": types.Integer(1), "income": types.Decimal(200)},
	)
	return be
}
