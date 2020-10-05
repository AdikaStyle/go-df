package go_df

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAnd(t *testing.T) {
	row := testRow()

	assert.True(t, And(
		Eq("int8", int8(12)),
		Eq("int16", int16(22)),
	)(row))

	assert.False(t, And(
		Eq("int8", int8(13)),
		Eq("int16", int16(22)),
	)(row))

	assert.False(t, And(
		Eq("int8", int8(12)),
		Eq("int16", int16(23)),
	)(row))

	assert.False(t, And(
		Eq("int8", int8(13)),
		Eq("int16", int16(23)),
	)(row))
}

func TestOr(t *testing.T) {
	row := testRow()

	assert.True(t, Or(
		Eq("int8", int8(12)),
		Eq("int16", int16(22)),
	)(row))

	assert.True(t, Or(
		Eq("int8", int8(13)),
		Eq("int16", int16(22)),
	)(row))

	assert.True(t, Or(
		Eq("int8", int8(12)),
		Eq("int16", int16(23)),
	)(row))

	assert.False(t, Or(
		Eq("int8", int8(13)),
		Eq("int16", int16(23)),
	)(row))
}

func TestNot(t *testing.T) {
	row := testRow()

	assert.True(t, Not(Eq("int8", int8(13)))(row))
	assert.False(t, Not(Eq("int8", int8(12)))(row))
}
