package dataframe

import (
	"github.com/AdikaStyle/go-df/backend"
	"github.com/AdikaStyle/go-df/conds"
)

type defaultJoinable struct {
	df Dataframe
}

func newDefaultJoinable(df Dataframe) *defaultJoinable {
	return &defaultJoinable{df: df}
}

func (this *defaultJoinable) LeftJoin(with Dataframe, on conds.JoinCondition) Dataframe {
	newHeaders := combineHeaders(this.df, with)
	joinedDf := this.df.constructNew(newHeaders)
	this.df.VisitRows(func(leftId int, left backend.Row) {
		added := false
		with.VisitRows(func(rightId int, right backend.Row) {
			if on(left, right) {
				row := combineRows(left, right, false)
				joinedDf.getBackend().AppendRow(row)
				added = true
			}
		})

		if !added {
			row := combineRows(left, with.getBackend().GetRow(0), true)
			joinedDf.getBackend().AppendRow(row)
		}
	})

	return joinedDf
}

func (this *defaultJoinable) RightJoin(with Dataframe, on conds.JoinCondition) Dataframe {
	return with.LeftJoin(this.df, on)
}

func (this *defaultJoinable) InnerJoin(with Dataframe, on conds.JoinCondition) Dataframe {
	newHeaders := combineHeaders(this.df, with)
	joinedDf := this.df.constructNew(newHeaders)
	this.df.VisitRows(func(leftId int, left backend.Row) {
		with.VisitRows(func(rightId int, right backend.Row) {
			if on(left, right) {
				joinedRow := combineRows(left, right, false)
				joinedDf.getBackend().AppendRow(joinedRow)
			}
		})
	})

	return joinedDf
}

func (this *defaultJoinable) OuterJoin(with Dataframe, on conds.JoinCondition) Dataframe {
	panic("unimplemented")
}

func combineHeaders(left Dataframe, right Dataframe) backend.Headers {
	var newHeaders backend.Headers
	dup := make(map[string]bool)

	for _, h := range left.GetHeaders() {
		newHeaders = append(newHeaders, h)
		dup[h.Name] = true
	}

	for _, h := range right.GetHeaders() {
		if _, found := dup[h.Name]; !found {
			newHeaders = append(newHeaders, h)
		}
	}

	return newHeaders
}

func combineRows(left backend.Row, right backend.Row, rightMissing bool) backend.Row {
	dup := make(map[string]bool)
	row := make(backend.Row)

	for k, v := range left {
		dup[k] = true
		row[k] = v
	}

	for k, v := range right {
		if _, found := dup[k]; found {
			continue
		}
		if rightMissing {
			row[k] = nil
		} else {
			row[k] = v
		}
	}

	return row
}
