package dataframe

import (
	"github.com/AdikaStyle/go-df/backend"
	"github.com/AdikaStyle/go-df/conds"
)

type defaultFilterable struct {
	df Dataframe
}

func newDefaultFilterable(dataframe Dataframe) *defaultFilterable {
	return &defaultFilterable{df: dataframe}
}

func (this *defaultFilterable) Select(columns ...string) Dataframe {
	selection := make(map[string]bool)
	for _, col := range columns {
		selection[col] = true
	}

	for _, h := range this.df.getBackend().GetHeaders() {
		if _, found := selection[h.Name]; !found {
			selection[h.Name] = false
		}
	}

	for k, keep := range selection {
		if !keep {
			this.df.getBackend().RemoveColumn(k)
		}
	}

	return this.df
}

func (this *defaultFilterable) Filter(cond conds.Condition) Dataframe {
	this.df.getBackend().ForEachRow(func(id int, row backend.Row) {
		if !cond(row) {
			this.df.getBackend().RemoveRows(id)
		}
	})

	return this.df
}
