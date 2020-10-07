package types

import (
	"fmt"
	"github.com/cstockton/go-conv"
	"time"
)

type DateTimeType time.Time

func (i DateTimeType) String() string {
	return time.Time(i).String()
}

func (i DateTimeType) Equals(other Type) bool {
	if i.Kind() == other.Kind() {
		otherVal, err := conv.Time(other)
		PanicOnError(err)
		return time.Time(i).Equal(otherVal)
	}
	return false
}

func (i DateTimeType) Compare(other Type) TypeComparision {
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

func (i DateTimeType) Cast(toPtr interface{}) {
	if err := conv.Infer(toPtr, i); err != nil {
		panic(err)
	}
}

func (i DateTimeType) Kind() TypeKind {
	return DateTime
}
