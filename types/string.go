package types

import (
	"fmt"
	"github.com/cstockton/go-conv"
)

type StringType string

func (i StringType) String() string {
	return string(i)
}

func (i StringType) Equals(other Type) bool {
	if i.Kind() == other.Kind() {
		return i == other
	}
	return false
}

func (i StringType) Compare(other Type) TypeComparision {
	if i.Kind() != other.Kind() {
		panic(fmt.Sprintf(
			"couldn't compare between different kind of types where left is: %s(%v) and right is: %s(%v)",
			i.Kind(), i, other.Kind(), other),
		)
	}

	otherVal, err := conv.String(other)
	PanicOnError(err)
	if string(i) == otherVal {
		return Equals
	} else if string(i) > otherVal {
		return LeftIsBigger
	} else {
		return RightIsBigger
	}
}

func (i StringType) Cast(toPtr interface{}) {
	if err := conv.Infer(toPtr, i); err != nil {
		panic(err)
	}
}

func (i StringType) Kind() TypeKind {
	return String
}
