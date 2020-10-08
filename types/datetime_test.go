package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var datetime1999 = Datetime(time.Date(1999, 1, 1, 10, 0, 0, 0, time.UTC))
var datetime1999_1 = Datetime(time.Date(1999, 1, 1, 10, 0, 0, 0, time.UTC))
var datetime2010 = Datetime(time.Date(2010, 1, 1, 10, 0, 0, 0, time.UTC))

func TestDateTimeType_String(t *testing.T) {
	assert.EqualValues(t, "1999-01-01 10:00:00 +0000 UTC", datetime1999.String())
	assert.EqualValues(t, "2010-01-01 10:00:00 +0000 UTC", datetime2010.String())
}

func TestDateTimeType_Kind(t *testing.T) {
	assert.EqualValues(t, KindDatetime, datetime1999.Kind())
	assert.EqualValues(t, KindDatetime, datetime2010.Kind())
}

func TestDateTimeType_Equals(t *testing.T) {
	assert.True(t, datetime1999.Equals(datetime1999_1))
	assert.False(t, datetime1999.Equals(datetime2010))
	assert.False(t, datetime2010.Equals(datetime1999))
}

func TestDateTimeType_Compare(t *testing.T) {
	assert.EqualValues(t, Equals, datetime1999.Compare(datetime1999_1))
	assert.EqualValues(t, LeftIsBigger, datetime2010.Compare(datetime1999))
	assert.EqualValues(t, RightIsBigger, datetime1999.Compare(datetime2010))
}

func TestDateTimeType_Cast(t *testing.T) {
	var str string
	datetime2010.Cast(&str)
	assert.EqualValues(t, "2010-01-01 10:00:00 +0000 UTC", str)
}
