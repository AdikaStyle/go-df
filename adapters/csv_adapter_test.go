package adapters

import (
	"fmt"
	"github.com/AdikaStyle/go-df/aggs"
	"github.com/AdikaStyle/go-df/backend"
	"github.com/AdikaStyle/go-df/conds"
	df "github.com/AdikaStyle/go-df/dataframe"
	"github.com/AdikaStyle/go-df/types"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
)

func TestLoadCSV(t *testing.T) {
	od := LoadCSV("../examples/order_details.csv", '|')
	assert.NotNil(t, od)

	products := LoadCSV("../examples/products.csv", '|')
	assert.NotNil(t, products)

	products.Select("ProductID", "ProductName")

	aggregated := od.
		Group(df.By{"ProductID"}, df.Aggs{"UnitPrice": aggs.Sum(), "Quantity": aggs.Sum()}).
		OrderBy("ProductID", backend.Asc).
		InnerJoin(products, conds.On("ProductID")).
		Filter(conds.Gte("UnitPrice", types.Decimal(1500)))

	aggregated.Print(os.Stdout)

	path := filepath.Join(os.TempDir(), fmt.Sprintf("go-df-%d.csv", rand.Int()))
	fmt.Printf("output saved to: %s\n", path)
	WriteCSV(path, aggregated, '|')

	aggregated2 := LoadCSV(path, '|')

	rows := mapOfProductId(aggregated2)

	aggregated.VisitRows(func(id int, row backend.Row) {
		other, found := rows[row["ProductID"].String()]
		assert.True(t, found)
		assert.NotNil(t, other)
		assert.EqualValues(t, row, other)
	})
}

func mapOfProductId(aggregated2 df.Dataframe) map[string]backend.Row {
	rows := make(map[string]backend.Row)
	aggregated2.VisitRows(func(id int, row backend.Row) {
		rcopy := make(backend.Row)
		for k, v := range row {
			rcopy[k] = v
		}
		rows[row["ProductID"].String()] = rcopy
	})
	return rows
}
