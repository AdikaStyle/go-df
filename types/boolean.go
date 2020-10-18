package types

import (
	"fmt"
	"github.com/cstockton/go-conv"
)

type Boolean bool

func (b Boolean) String() string {
	if b {
		return "true"
	} else {
		return "false"
	}
}

func (b Boolean) Equals(other TypedValue) bool {
	if b.Kind() == other.Kind() {
		return b == other
	}
	return false
}

func (b Boolean) Compare(other TypedValue) TypeComparision {
	if b.Kind() != other.Kind() {
		panic(fmt.Sprintf(
			"couldn't compare between different kind of types where left is: %s(%v) and right is: %s(%v)",
			b.Kind(), b, other.Kind(), other),
		)
	}

	var otherVal bool
	other.Cast(&otherVal)

	if b == other {
		return Equals
	} else if b == true && otherVal == false {
		return LeftIsBigger
	} else {
		return RightIsBigger
	}
}

func (b Boolean) Cast(toPtr interface{}) {
	if err := conv.Infer(toPtr, b); err != nil {
		panic(err)
	}
}

func (b Boolean) Kind() TypeKind {
	return KindBoolean
}

func (b Boolean) Ptr() TypedValue {
	return &b
}

func (b Boolean) NativeType() interface{} {
	return bool(b)
}

func (b Boolean) Precedence() int {
	return 0
}
