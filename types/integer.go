package types

import (
	"fmt"
	"github.com/cstockton/go-conv"
)

type IntegerType int64

func (i IntegerType) String() string {
	str, err := conv.String(int64(i))
	if err != nil {
		panic(err)
	}
	return str
}

func (i IntegerType) Equals(other Type) bool {
	if i.Kind() == other.Kind() {
		return i == other
	}
	return false
}

func (i IntegerType) Compare(other Type) TypeComparision {
	if i.Kind() != other.Kind() {
		panic(fmt.Sprintf(
			"couldn't compare between different kind of types where left is: %s(%v) and right is: %s(%v)",
			i.Kind(), i, other.Kind(), other),
		)
	}

	otherVal, err := conv.Int64(other)
	PanicOnError(err)
	if int64(i) == otherVal {
		return Equals
	} else if int64(i) > otherVal {
		return LeftIsBigger
	} else {
		return RightIsBigger
	}
}

func (i IntegerType) Cast(toPtr interface{}) {
	if err := conv.Infer(toPtr, i); err != nil {
		panic(err)
	}
}

func (i IntegerType) Kind() TypeKind {
	return Integer
}
