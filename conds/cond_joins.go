package conds

import "github.com/AdikaStyle/go-df/backend"

func On(column string, columns ...string) JoinCondition {
	return func(left backend.Row, right backend.Row) bool {
		var out = left[column].Equals(right[column])
		for _, col := range columns {
			out = out && (left[col] == right[col])
		}
		return out
	}
}
