package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const one = IntegerType(1)
const one_1 = IntegerType(1)
const two = IntegerType(2)

func TestIntegerType_String(t *testing.T) {
	assert.EqualValues(t, "1", one.String())
	assert.EqualValues(t, "2", two.String())
}

func TestIntegerType_Kind(t *testing.T) {
	assert.EqualValues(t, Integer, one.Kind())
	assert.EqualValues(t, Integer, two.Kind())
}

func TestIntegerType_Equals(t *testing.T) {
	assert.True(t, one.Equals(one_1))
	assert.False(t, one.Equals(two))
	assert.False(t, two.Equals(one))
}

func TestIntegerType_Compare(t *testing.T) {
	assert.EqualValues(t, Equals, one.Compare(one_1))
	assert.EqualValues(t, LeftIsBigger, two.Compare(one))
	assert.EqualValues(t, RightIsBigger, one.Compare(two))
}

func TestIntegerType_Cast(t *testing.T) {
	var str string
	two.Cast(&str)
	assert.EqualValues(t, "2", str)

	var decimal float64
	two.Cast(&decimal)
	assert.EqualValues(t, 2, decimal)
}
