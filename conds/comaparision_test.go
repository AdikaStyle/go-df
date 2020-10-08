package conds

import (
	"github.com/AdikaStyle/go-df/backend"
	"github.com/AdikaStyle/go-df/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEq(t *testing.T) {
	row := testRow()

	assert.True(t, Eq("bool", types.Boolean(true))(row))
	assert.False(t, Eq("bool", types.Boolean(false))(row))
	assert.True(t, Eq("int64", types.Integer(2200000))(row))
	assert.False(t, Eq("int64", types.Integer(1))(row))
	assert.True(t, Eq("float32", types.Decimal(22.33))(row))
	assert.False(t, Eq("float32", types.Decimal(22.34))(row))
	assert.True(t, Eq("float64", types.Decimal(40.2))(row))
	assert.False(t, Eq("float64", types.Decimal(40.3))(row))
	assert.True(t, Eq("string", types.String("Hello World"))(row))
	assert.False(t, Eq("string", types.String("Hello Worldq"))(row))
}

func TestGt(t *testing.T) {
	row := testRow()

	assert.True(t, Gt("int64", types.Integer(10))(row))
	assert.True(t, Gte("int64", types.Integer(10))(row))
	assert.True(t, Gte("int64", types.Integer(2200000))(row))
	assert.False(t, Gte("int64", types.Integer(2200001))(row))
}

func TestLt(t *testing.T) {
	row := testRow()

	assert.False(t, Lt("int64", types.Integer(10))(row))
	assert.True(t, Lt("int8", types.Integer(23))(row))
	assert.False(t, Lte("int64", types.Integer(10))(row))
	assert.True(t, Lte("int64", types.Integer(2200000))(row))
	assert.True(t, Lte("int64", types.Integer(2200001))(row))
}

func TestIn(t *testing.T) {
	row := testRow()

	assert.True(t, In("int8", types.Integer(1), types.Integer(2), types.Integer(12))(row))
	assert.True(t, In("int8", types.Integer(80), types.Integer(12), types.Integer(-12))(row))
	assert.False(t, In("int8", types.Integer(1), types.Integer(2), types.Integer(3))(row))
}

func testRow() backend.Row {
	row := make(backend.Row)
	row["bool"] = types.Boolean(true)
	row["int8"] = types.Integer(12)
	row["int16"] = types.Integer(22)
	row["int32"] = types.Integer(1002)
	row["int64"] = types.Integer(2200000)
	row["float32"] = types.Decimal(22.33)
	row["float64"] = types.Decimal((40.2))
	row["string"] = types.String("Hello World")
	return row
}
