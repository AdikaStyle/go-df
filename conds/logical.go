package conds

import "github.com/AdikaStyle/go-df/backend"

func Not(val Condition) Condition {
	return func(row backend.Row) bool {
		return !val(row)
	}
}

func And(left Condition, right Condition) Condition {
	return func(row backend.Row) bool {
		return left(row) && right(row)
	}
}

func Or(left Condition, right Condition) Condition {
	return func(row backend.Row) bool {
		return left(row) || right(row)
	}
}
