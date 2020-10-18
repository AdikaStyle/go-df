package types

import (
	"fmt"
	"github.com/cstockton/go-conv"
)

type Integer int64

func (i Integer) String() string {
	str, err := conv.String(int64(i))
	if err != nil {
		panic(err)
	}
	return str
}

func (i Integer) Equals(other TypedValue) bool {
	if i.Kind() == other.Kind() {
		return i == other
	}
	return false
}

func (i Integer) Compare(other TypedValue) TypeComparision {
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

func (i Integer) Cast(toPtr interface{}) {
	if err := conv.Infer(toPtr, i); err != nil {
		panic(err)
	}
}

func (i Integer) Kind() TypeKind {
	return KindInteger
}

func (i Integer) Ptr() TypedValue {
	return &i
}

func (i Integer) NativeType() interface{} {
	return int64(i)
}

func (i Integer) Precedence() int {
	return 1
}
