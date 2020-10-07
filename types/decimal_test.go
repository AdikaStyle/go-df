package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const pi = DecimalType(3.14)
const pi_1 = DecimalType(3.14)
const e = DecimalType(2.71)

func TestDecimalType_String(t *testing.T) {
	assert.EqualValues(t, "3.14", pi.String())
	assert.EqualValues(t, "2.71", e.String())
}

func TestDecimalType_Kind(t *testing.T) {
	assert.EqualValues(t, Decimal, pi.Kind())
	assert.EqualValues(t, Decimal, e.Kind())
}

func TestDecimalType_Equals(t *testing.T) {
	assert.True(t, pi.Equals(pi_1))
	assert.False(t, pi.Equals(e))
	assert.False(t, e.Equals(pi))
}

func TestDecimalType_Compare(t *testing.T) {
	assert.EqualValues(t, Equals, pi.Compare(pi_1))
	assert.EqualValues(t, LeftIsBigger, pi.Compare(e))
	assert.EqualValues(t, RightIsBigger, e.Compare(pi))
}

func TestDecimalType_Cast(t *testing.T) {
	var str string
	pi.Cast(&str)
	assert.EqualValues(t, "3.14", str)

	var integer int64
	pi.Cast(&integer)
	assert.EqualValues(t, 3, integer)
}
