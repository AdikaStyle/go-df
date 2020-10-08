package conds

import (
	"github.com/AdikaStyle/go-df/backend"
	"github.com/AdikaStyle/go-df/types"
)

func Eq(columnName string, value types.TypedValue) Condition {
	return func(row backend.Row) bool {
		return value.Equals(row[columnName])
	}
}

func NotEq(columnName string, value types.TypedValue) Condition {
	return func(row backend.Row) bool {
		return !value.Equals(row[columnName])
	}
}

func Gt(columnName string, value types.TypedValue) Condition {
	return func(row backend.Row) bool {
		left := row[columnName]
		return left.Compare(value) == types.LeftIsBigger
	}
}

func Gte(columnName string, value types.TypedValue) Condition {
	return func(row backend.Row) bool {
		left := row[columnName]

		cmp := left.Compare(value)
		return cmp == types.Equals || cmp == types.LeftIsBigger
	}
}

func Lt(columnName string, value types.TypedValue) Condition {
	return func(row backend.Row) bool {
		left := row[columnName]
		return left.Compare(value) == types.RightIsBigger
	}
}

func Lte(columnName string, value types.TypedValue) Condition {
	return func(row backend.Row) bool {
		left := row[columnName]

		cmp := left.Compare(value)
		return cmp == types.Equals || cmp == types.RightIsBigger
	}
}

func In(columnName string, inValues ...types.TypedValue) Condition {
	return func(row backend.Row) bool {
		for _, val := range inValues {
			if val.Equals(row[columnName]) {
				return true
			}
		}
		return false
	}
}
