package dataframe

import (
	"github.com/AdikaStyle/go-df/aggs"
	"github.com/AdikaStyle/go-df/backend"
	"github.com/AdikaStyle/go-df/conds"
	"github.com/AdikaStyle/go-df/types"
	"io"
)

func New(be backend.Backend) Dataframe {
	return newColumnarDataframe(be)
}

type Dataframe interface {
	Filterable
	Transformable
	Aggregatable
	Joinable

	VisitRows(visitor backend.RowVisitor) Dataframe
	VisitColumn(visitor backend.ColumnVisitor)
	GetRowCount() int
	GetHeaders() backend.Headers

	constructNew(headers backend.Headers) Dataframe
	getBackend() backend.Backend

	Printable
}

type Filterable interface {
	Select(columns ...string) Dataframe
	Filter(cond conds.Condition) Dataframe
}

type Transformable interface {
	UpdateColumn(name string, fn backend.MutateFunction) Dataframe
	AddColumn(name string, kind types.TypeKind, fn backend.MutateFunction) Dataframe
	Concat(with Dataframe) Dataframe
	Split(cond conds.Condition) (onTrue Dataframe, onFalse Dataframe)
	CastColumn(name string, to types.TypeKind) Dataframe
	Apply(name string, fn backend.ApplyFunction) Dataframe
}

type Aggregatable interface {
	Group(columns []string, aggregations map[string]aggs.Aggregation) Dataframe
	OrderBy(columns string, order backend.Ordering) Dataframe
}

type Joinable interface {
	LeftJoin(with Dataframe, on conds.JoinCondition) Dataframe
	RightJoin(with Dataframe, on conds.JoinCondition) Dataframe
	InnerJoin(with Dataframe, on conds.JoinCondition) Dataframe
	OuterJoin(with Dataframe, on conds.JoinCondition) Dataframe
}

type Printable interface {
	Print(w io.Writer) Dataframe
	Describe(w io.Writer) Dataframe
}

type By []string
type Aggs map[string]aggs.Aggregation
