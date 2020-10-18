package types

import (
	"fmt"
	"github.com/cstockton/go-conv"
)

type String string

func (i String) String() string {
	return string(i)
}

func (i String) Equals(other TypedValue) bool {
	if i.Kind() == other.Kind() {
		return i == other
	}
	return false
}

func (i String) Compare(other TypedValue) TypeComparision {
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

func (i String) Cast(toPtr interface{}) {
	if err := conv.Infer(toPtr, i); err != nil {
		panic(err)
	}
}

func (i String) Kind() TypeKind {
	return KindString
}

func (i String) Ptr() TypedValue {
	return &i
}

func (i String) NativeType() interface{} {
	return string(i)
}

func (i String) Precedence() int {
	return 4
}
