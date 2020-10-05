package go_df

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEq(t *testing.T) {
	row := testRow()

	assert.True(t, Eq("bool", true)(row))
	assert.False(t, Eq("bool", false)(row))
	assert.True(t, Eq("int8", int8(12))(row))
	assert.False(t, Eq("int8", int8(18))(row))
	assert.True(t, Eq("int16", int16(22))(row))
	assert.False(t, Eq("int16", int16(-1))(row))
	assert.True(t, Eq("int32", int32(1002))(row))
	assert.False(t, Eq("int32", int32(100233))(row))
	assert.True(t, Eq("int64", int64(2200000))(row))
	assert.False(t, Eq("int64", int64(1))(row))
	assert.True(t, Eq("float32", float32(22.33))(row))
	assert.False(t, Eq("float32", float32(22.34))(row))
	assert.True(t, Eq("float64", float64(40.2))(row))
	assert.False(t, Eq("float64", float64(40.3))(row))
	assert.True(t, Eq("string", "Hello World")(row))
	assert.False(t, Eq("string", "Hello Worldq")(row))
}

func TestGt(t *testing.T) {
	row := testRow()

	assert.True(t, Gt("int64", 10)(row))
	assert.False(t, Gt("int8", 23)(row))
	assert.True(t, Gte("int64", 10)(row))
	assert.True(t, Gte("int64", 2200000)(row))
	assert.False(t, Gte("int64", 2200001)(row))
}

func TestLt(t *testing.T) {
	row := testRow()

	assert.False(t, Lt("int64", 10)(row))
	assert.True(t, Lt("int8", 23)(row))
	assert.False(t, Lte("int64", 10)(row))
	assert.True(t, Lte("int64", 2200000)(row))
	assert.True(t, Lte("int64", 2200001)(row))
}

func TestIn(t *testing.T) {
	row := testRow()

	assert.True(t, In("int8", int8(1), int8(2), int8(12))(row))
	assert.True(t, In("int8", int8(80), int8(12), int8(-12))(row))
	assert.False(t, In("int8", int8(1), int8(2), int8(3))(row))
}

func testRow() Row {
	row := make(Row)
	row["bool"] = true
	row["int8"] = int8(12)
	row["int16"] = int16(22)
	row["int32"] = int32(1002)
	row["int64"] = int64(2200000)
	row["float32"] = float32(22.33)
	row["float64"] = float64(40.2)
	row["string"] = "Hello World"
	return row
}
