package dataframe

import (
	"fmt"
	"github.com/AdikaStyle/go-df/aggs"
	"github.com/AdikaStyle/go-df/backend"
	"log"
	"strings"
)

type defaultAggregatable struct {
	df Dataframe
}

func newDefaultAggregatable(df Dataframe) *defaultAggregatable {
	return &defaultAggregatable{df: df}
}

type group map[string]backend.Columns

func (this *defaultAggregatable) Group(columns []string, aggregations map[string]aggs.Aggregation) Dataframe {
	groups := make(group)

	this.df.VisitRows(func(id int, row backend.Row) {
		groupKey := encodeGroupKey(columns, row)
		if groups[groupKey] == nil {
			groups[groupKey] = make(backend.Columns)
		}

		for k, v := range row {
			groups[groupKey][k] = append(groups[groupKey][k], backend.Cell{TypedValue: v})
		}
	})

	outCols := make(backend.Columns)
	for _, v := range groups {
		for col, agg := range aggregations {
			aggVal := agg(v[col])
			outCols[col] = append(outCols[col], backend.Cell{TypedValue: aggVal})
		}

		for _, c := range columns {
			if v[c] == nil {
				log.Fatalf("failed to find column named: %s in dataframe", c)
			}
			outCols[c] = append(outCols[c], backend.Cell{TypedValue: aggs.First()(v[c])})
		}
	}

	out := this.df.constructNew(outCols.GetHeaders())
	out.getBackend().SetColumns(outCols)
	return out
}

func (this *defaultAggregatable) OrderBy(columns string, order backend.Ordering) Dataframe {
	this.df.getBackend().SortByColumn(columns, order)
	return this.df
}

func encodeGroupKey(cols []string, row backend.Row) string {
	b := strings.Builder{}
	for _, c := range cols {
		b.WriteString(fmt.Sprintf("%v ", row[c]))
	}
	return b.String()
}
