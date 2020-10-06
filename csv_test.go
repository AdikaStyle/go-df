package go_df

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestLoadCsv(t *testing.T) {
	od, err := LoadCSV("examples/order_details.csv", '|')
	assert.Nil(t, err)
	assert.NotNil(t, od)

	products, err := LoadCSV("examples/products.csv", '|')
	assert.Nil(t, err)
	assert.NotNil(t, products)

	od = od.GroupBy(Cols{"ProductID"}, Aggs{"UnitPrice": Sum(), "Quantity": Sum()})
	od.Print()
	od.OrderBy("ProductID", -1).Print()

	odPath := filepath.Join(os.TempDir(), "b.csv")
	println(odPath)
	WriteCSV(odPath, '|', od)
}
