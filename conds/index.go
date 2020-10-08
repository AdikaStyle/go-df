package conds

import "github.com/AdikaStyle/go-df/backend"

type Condition func(row backend.Row) bool

type JoinCondition func(left backend.Row, right backend.Row) bool
