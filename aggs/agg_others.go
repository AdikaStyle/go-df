package aggs

import (
	"github.com/AdikaStyle/go-df/backend"
	"github.com/AdikaStyle/go-df/types"
)

func Count() Aggregation {
	return func(column backend.Column) types.TypedValue {
		return types.Integer(len(column))
	}
}

func CountDistinct() Aggregation {
	return func(column backend.Column) types.TypedValue {
		dup := make(map[types.TypedValue]bool)
		for idx := range column {
			dup[column[idx]] = true
		}
		return types.Integer(len(dup))
	}
}

func First() Aggregation {
	return func(column backend.Column) types.TypedValue {
		return column[0]
	}
}

func Last() Aggregation {
	return func(column backend.Column) types.TypedValue {
		return column[len(column)-1]
	}
}
