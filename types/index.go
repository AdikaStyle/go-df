package types

import "fmt"

type TypeComparision int

const (
	Equals        TypeComparision = 0
	LeftIsBigger  TypeComparision = 1
	RightIsBigger TypeComparision = -1
)

type TypeKind string

const (
	KindBoolean  TypeKind = "Boolean"
	KindInteger  TypeKind = "Integer"
	KindDecimal  TypeKind = "Decimal"
	KindString   TypeKind = "String"
	KindDatetime TypeKind = "Datetime"

	KindMissing TypeKind = "KindMissing"
	KindUnknown TypeKind = "KindUnknown"
)

type TypedValue interface {
	fmt.Stringer
	Equals(other TypedValue) bool
	Compare(other TypedValue) TypeComparision
	Cast(toPtr interface{})
	Kind() TypeKind
	Ptr() TypedValue
	NativeType() interface{}
	Precedence() int
}
