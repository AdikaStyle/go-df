package dataframe

import (
	"github.com/AdikaStyle/go-df/backend"
	"github.com/AdikaStyle/go-df/conds"
	"github.com/AdikaStyle/go-df/types"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestDefaultFilterable_Select(t *testing.T) {
	df := newColumnarDataframe(filterBackend())
	df.getBackend().AppendRow(
		backend.Row{"name": types.String("Abraham"), "age": types.Integer(102), "brit": types.Boolean(true)},
		backend.Row{"name": types.String("Yitzack"), "age": types.Integer(203), "brit": types.Boolean(true)},
		backend.Row{"name": types.String("Yaakov"), "age": types.Integer(99), "brit": types.Boolean(true)},
	)

	df = df.Select("name", "brit")
	headers := df.GetHeaders()

	assert.EqualValues(t, 2, len(headers))
	assert.EqualValues(t, "name", headers[0].Name)
	assert.EqualValues(t, "brit", headers[1].Name)

	df.Print(os.Stdout)
}

func TestDefaultFilterable_Filter(t *testing.T) {
	df := newColumnarDataframe(filterBackend())
	df.getBackend().AppendRow(
		backend.Row{"name": types.String("Abraham"), "age": types.Integer(102), "brit": types.Boolean(true)},
		backend.Row{"name": types.String("Yitzack"), "age": types.Integer(203), "brit": types.Boolean(true)},
		backend.Row{"name": types.String("Yaakov"), "age": types.Integer(99), "brit": types.Boolean(true)},
	)

	df = df.Filter(conds.Gt("age", types.Integer(100)))

	assert.EqualValues(t, 2, df.GetRowCount())
	df.VisitRows(func(id int, row backend.Row) {
		assert.NotEqualValues(t, "Yaakov", row["name"])
	})

	df.Print(os.Stdout)
}

func filterBackend() backend.Backend {
	headers := backend.Headers{
		{Name: "name", Kind: types.KindString},
		{Name: "age", Kind: types.KindInteger},
		{Name: "brit", Kind: types.KindBoolean},
	}
	be := backend.NewColumnarBackend(headers)
	return be
}
