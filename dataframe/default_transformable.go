package dataframe

import (
	"github.com/AdikaStyle/go-df/backend"
	"github.com/AdikaStyle/go-df/conds"
)

type defaultTransformable struct {
	df Dataframe
}

func newDefaultTransformable(df Dataframe) *defaultTransformable {
	return &defaultTransformable{df: df}
}

func (this *defaultTransformable) Concat(with Dataframe) Dataframe {
	with.VisitRows(func(id int, row backend.Row) {
		this.df.getBackend().AppendRow(row)
	})
	return this.df
}

func (this *defaultTransformable) Split(cond conds.Condition) (onTrue Dataframe, onFalse Dataframe) {
	onTrue = this.df.constructNew(this.df.GetHeaders())
	onFalse = this.df.constructNew(this.df.GetHeaders())

	this.df.VisitRows(func(id int, row backend.Row) {
		if cond(row) {
			onTrue.getBackend().AppendRow(row)
		} else {
			onFalse.getBackend().AppendRow(row)
		}
	})

	return onTrue, onFalse
}
