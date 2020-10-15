package conds

import (
	"github.com/AdikaStyle/go-df/backend"
	"strings"
)

func On(column string, moreColumns ...string) JoinCondition {
	return newOnJoinCondition(append(moreColumns, column))
}

type onJoinCondition struct {
	columns []string
}

func newOnJoinCondition(columns []string) *onJoinCondition {
	return &onJoinCondition{columns: columns}
}

func (o onJoinCondition) Match(left backend.Row, right backend.Row) bool {
	var out = true
	for _, col := range o.columns {
		out = out && (left[col] == right[col])
	}
	return out
}

func (o onJoinCondition) ColumnsHint(row backend.Row) string {
	sb := &strings.Builder{}
	for _, col := range o.columns {
		sb.WriteString(row[col].String())
	}
	return sb.String()
}
