package types

import (
	"fmt"
	"github.com/cstockton/go-conv"
	"time"
)

type Datetime time.Time

func (i Datetime) String() string {
	return time.Time(i).String()
}

func (i Datetime) Equals(other TypedValue) bool {
	if i.Kind() == other.Kind() {
		otherVal, err := conv.Time(other)
		PanicOnError(err)
		return time.Time(i).Equal(otherVal)
	}
	return false
}

func (i Datetime) Compare(other TypedValue) TypeComparision {
	if i.Kind() != other.Kind() {
		panic(fmt.Sprintf(
			"couldn't compare between different kind of types where left is: %s(%v) and right is: %s(%v)",
			i.Kind(), i, other.Kind(), other),
		)
	}

	otherVal, err := conv.Time(other)
	PanicOnError(err)
	if time.Time(i).Equal(otherVal) {
		return Equals
	} else if time.Time(i).After(otherVal) {
		return LeftIsBigger
	} else {
		return RightIsBigger
	}
}

func (i Datetime) Cast(toPtr interface{}) {
	if err := conv.Infer(toPtr, i); err != nil {
		panic(err)
	}
}

func (i Datetime) Kind() TypeKind {
	return KindDatetime
}

func (i Datetime) Ptr() TypedValue {
	return &i
}

func (i Datetime) NativeType() interface{} {
	return time.Time(i)
}

func (i Datetime) Precedence() int {
	return 3
}
