package conds

import "github.com/AdikaStyle/go-df/backend"

type Condition func(row backend.Row) bool

type JoinCondition interface {
	Match(left backend.Row, right backend.Row) bool
	ColumnsHint(row backend.Row) string
}
