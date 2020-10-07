package types

import (
	"fmt"
	"github.com/cstockton/go-conv"
)

type BooleanType bool

func (b BooleanType) String() string {
	if b {
		return "true"
	} else {
		return "false"
	}
}

func (b BooleanType) Equals(other Type) bool {
	if b.Kind() == other.Kind() {
		return b == other
	}
	return false
}

func (b BooleanType) Compare(other Type) TypeComparision {
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

func (b BooleanType) Cast(toPtr interface{}) {
	if err := conv.Infer(toPtr, b); err != nil {
		panic(err)
	}
}

func (b BooleanType) Kind() TypeKind {
	return Boolean
}
