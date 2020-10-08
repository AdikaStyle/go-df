package conds

import (
	"github.com/AdikaStyle/go-df/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAnd(t *testing.T) {
	row := testRow()

	assert.True(t, And(
		Eq("int8", types.Integer(12)),
		Eq("int16", types.Integer(22)),
	)(row))

	assert.False(t, And(
		Eq("int8", types.Integer(13)),
		Eq("int16", types.Integer(22)),
	)(row))

	assert.False(t, And(
		Eq("int8", types.Integer(12)),
		Eq("int16", types.Integer(23)),
	)(row))

	assert.False(t, And(
		Eq("int8", types.Integer(13)),
		Eq("int16", types.Integer(23)),
	)(row))
}

func TestOr(t *testing.T) {
	row := testRow()

	assert.True(t, Or(
		Eq("int8", types.Integer(12)),
		Eq("int16", types.Integer(22)),
	)(row))

	assert.True(t, Or(
		Eq("int8", types.Integer(13)),
		Eq("int16", types.Integer(22)),
	)(row))

	assert.True(t, Or(
		Eq("int8", types.Integer(12)),
		Eq("int16", types.Integer(23)),
	)(row))

	assert.False(t, Or(
		Eq("int8", types.Integer(13)),
		Eq("int16", types.Integer(23)),
	)(row))
}

func TestNot(t *testing.T) {
	row := testRow()

	assert.True(t, Not(Eq("int8", types.Integer(13)))(row))
	assert.False(t, Not(Eq("int8", types.Integer(12)))(row))
}
