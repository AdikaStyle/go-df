package dataframe

import (
	"github.com/AdikaStyle/go-df/backend"
)

type columnarDataframe struct {
	backend backend.Backend
	Filterable
	Transformable
	Aggregatable
	Joinable
	Printable
}

func newColumnarDataframe(backend backend.Backend) Dataframe {
	df := &columnarDataframe{backend: backend}
	df.Filterable = newDefaultFilterable(df)
	df.Transformable = newDefaultTransformable(df)
	df.Aggregatable = newDefaultAggregatable(df)
	df.Joinable = newDefaultJoinable(df)
	df.Printable = newDefaultPrintable(df)
	return df
}

func (this *columnarDataframe) VisitRows(visitor backend.RowVisitor) Dataframe {
	this.backend.ForEachRow(visitor)
	return this
}

func (this *columnarDataframe) VisitColumn(visitor backend.ColumnVisitor) {
	this.backend.ForEachColumn(visitor)
}

func (this *columnarDataframe) GetRowCount() int {
	return this.backend.GetRowCount()
}

func (this *columnarDataframe) GetHeaders() backend.Headers {
	return this.backend.GetHeaders()
}

func (this *columnarDataframe) getBackend() backend.Backend {
	return this.backend
}

func (this *columnarDataframe) constructNew(headers backend.Headers) Dataframe {
	return newColumnarDataframe(this.backend.ConstructNew(headers))
}
