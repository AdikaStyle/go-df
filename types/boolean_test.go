package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const trueType = Boolean(true)
const falseType = Boolean(false)

func TestBooleanType_String(t *testing.T) {
	assert.EqualValues(t, "true", trueType.String())
	assert.EqualValues(t, "false", falseType.String())
}

func TestBooleanType_Kind(t *testing.T) {
	assert.EqualValues(t, KindBoolean, trueType.Kind())
	assert.EqualValues(t, KindBoolean, falseType.Kind())
}

func TestBooleanType_Equals(t *testing.T) {
	assert.True(t, trueType.Equals(trueType))
	assert.True(t, falseType.Equals(falseType))
	assert.False(t, trueType.Equals(falseType))
	assert.False(t, falseType.Equals(trueType))
}

func TestBooleanType_Compare(t *testing.T) {
	assert.EqualValues(t, Equals, trueType.Compare(trueType))
	assert.EqualValues(t, LeftIsBigger, trueType.Compare(falseType))
	assert.EqualValues(t, RightIsBigger, falseType.Compare(trueType))
}

func TestBooleanType_Cast(t *testing.T) {
	var str string
	trueType.Cast(&str)
	assert.EqualValues(t, "true", str)

	var integer int64
	trueType.Cast(&integer)
	assert.EqualValues(t, 1, integer)

	var decimal float64
	trueType.Cast(&decimal)
	assert.EqualValues(t, 1, decimal)
}
