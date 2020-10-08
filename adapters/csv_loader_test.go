package adapters

import (
	"github.com/AdikaStyle/go-df/aggs"
	"github.com/AdikaStyle/go-df/backend"
	"github.com/AdikaStyle/go-df/conds"
	df "github.com/AdikaStyle/go-df/dataframe"
	"github.com/AdikaStyle/go-df/types"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLoadCSV(t *testing.T) {
	od := LoadCSV("../examples/order_details.csv", '|')
	assert.NotNil(t, od)

	products := LoadCSV("../examples/products.csv", '|')
	assert.NotNil(t, products)

	products.Select("ProductID", "ProductName")

	od.
		Group(df.By{"ProductID"}, df.Aggs{"UnitPrice": aggs.Sum(), "Quantity": aggs.Sum()}).
		OrderBy("ProductID", backend.Asc).
		InnerJoin(products, conds.On("ProductID")).
		Filter(conds.Gte("UnitPrice", types.Decimal(600))).
		Print(os.Stdout)
}
