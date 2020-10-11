package backend

import "github.com/AdikaStyle/go-df/types"

func New(headers Headers) Backend {
	return NewColumnarBackend(headers)
}

type Backend interface {
	GetHeaders() Headers
	GetRowCount() int
	GetRow(id int) Row
	AppendRow(row ...Row)
	ForEachRow(visitor RowVisitor)
	ForEachColumn(visitor ColumnVisitor)
	AddColumn(name string, kind types.TypeKind, mutateFn MutateFunction)
	UpdateColumn(name string, mutateFn MutateFunction)
	RemoveRows(ids ...int)
	RemoveColumn(name string)
	RenameColumn(old string, new string)
	SortByColumn(column string, order Ordering)
	ConstructNew(headers Headers) Backend
	SetColumns(columns Columns)
}

type RowVisitor func(id int, row Row)

type ColumnVisitor func(header string, values Column)

type MutateFunction func(id int, row Row) types.TypedValue

type Row map[string]types.TypedValue

type Cell struct {
	types.TypedValue
	filtered bool
}

type Column []Cell

type Columns map[string]Column

type Headers []Header

type Header struct {
	Name    string
	Kind    types.TypeKind
	Visible bool
}

type Ordering bool

const (
	Asc  Ordering = true
	Desc Ordering = false
)
