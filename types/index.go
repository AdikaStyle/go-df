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
	Boolean  TypeKind = "Boolean"
	Integer  TypeKind = "Integer"
	Decimal  TypeKind = "Decimal"
	String   TypeKind = "String"
	DateTime TypeKind = "DateTime"

	Missing TypeKind = "Missing"
	Unknown TypeKind = "Unknown"
)

type Type interface {
	fmt.Stringer
	Equals(other Type) bool
	Compare(other Type) TypeComparision
	Cast(toPtr interface{})
	Kind() TypeKind
}
