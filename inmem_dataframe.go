package go_df

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"log"
	"math"
	"os"
	"strings"
)

type group map[string]map[string][]CValue

type inMemoryDataframe struct {
	cs *ColumnStore
}

func NewInMemoryDataframe(cs *ColumnStore) *inMemoryDataframe {
	return &inMemoryDataframe{cs: cs}
}

func (this *inMemoryDataframe) Select(columns ...string) DataFrame {
	selection := make(map[string]bool)
	for _, col := range columns {
		selection[col] = true
	}

	for _, header := range this.cs.headers {
		if header == "" {
			continue
		}
		if _, found := selection[header]; !found {
			this.cs.RemoveColumn(header)
		}
	}

	return this
}

func (this *inMemoryDataframe) Filter(cond Condition) DataFrame {
	var removals []int
	for i := 0; i < this.cs.GetRowsCount(); i++ {
		if !cond(this.cs.GetRow(i)) {
			removals = append(removals, i)
		}
	}
	this.cs.RemoveRows(removals...)
	return this
}

func (this *inMemoryDataframe) Concat(with DataFrame) DataFrame {
	if len(this.cs.store) != len(with.(*inMemoryDataframe).cs.store) {
		log.Fatal("in order to concat two data-frames their column count should match")
	}

	for k, v := range with.(*inMemoryDataframe).cs.store {
		if _, found := this.cs.store[k]; !found {
			log.Fatalf("error when trying to concat two data-frames, column %s doesn't exists in the base table", k)
		}

		this.cs.store[k] = append(this.cs.store[k], v...)
	}

	return this
}

func (this *inMemoryDataframe) Split(cond Condition) (onTrue DataFrame, onFalse DataFrame) {
	t := make(map[string][]CValue)
	f := make(map[string][]CValue)

	for i := 0; i < this.cs.GetRowsCount(); i++ {
		row := this.cs.GetRow(i)
		if cond(row) {
			for k, v := range row {
				t[k] = append(t[k], v)
			}
		} else {
			for k, v := range row {
				f[k] = append(f[k], v)
			}
		}
	}

	return NewInMemoryDataframe(NewColumnStore(t)), NewInMemoryDataframe(NewColumnStore(f))
}

func (this *inMemoryDataframe) GroupBy(columns []string, aggregations map[string]Aggregation) DataFrame {
	groups := make(group)

	for i := 0; i < this.GetRowCount(); i++ {
		row := this.cs.GetRow(i)
		groupKey := encodeGroupKey(columns, row)
		if groups[groupKey] == nil {
			groups[groupKey] = make(map[string][]CValue)
		}

		for k, v := range row {
			groups[groupKey][k] = append(groups[groupKey][k], v)
		}
	}

	store := make(map[string][]CValue)
	for _, v := range groups {
		for col, agg := range aggregations {
			aggVal := agg(v[col])
			store[col] = append(store[col], aggVal)
		}

		for _, c := range columns {
			store[c] = append(store[c], First()(v[c]))
		}
	}

	return NewInMemoryDataframe(NewColumnStore(store))
}

func (this *inMemoryDataframe) InnerJoin(with DataFrame, on JoinCondition) DataFrame {
	store := make(map[string][]CValue)
	for i := 0; i < this.GetRowCount(); i++ {
		for j := 0; j < with.GetRowCount(); j++ {
			left := this.cs.GetRow(i)
			right := with.(*inMemoryDataframe).cs.GetRow(j)
			if on(left, right) {
				combineRows(store, left, right, false)
			}
		}
	}

	return NewInMemoryDataframe(NewColumnStore(store))
}

func (this *inMemoryDataframe) LeftJoin(with DataFrame, on JoinCondition) DataFrame {
	store := make(map[string][]CValue)
	for i := 0; i < this.GetRowCount(); i++ {
		added := false
		left := this.cs.GetRow(i)
		for j := 0; j < with.GetRowCount(); j++ {
			right := with.(*inMemoryDataframe).cs.GetRow(j)
			if on(left, right) {
				combineRows(store, left, right, false)
				added = true
			}
		}

		if !added {
			combineRows(store, left, with.(*inMemoryDataframe).cs.GetRow(0), true)
		}
	}

	return NewInMemoryDataframe(NewColumnStore(store))
}

func (this *inMemoryDataframe) RightJoin(with DataFrame, on JoinCondition) DataFrame {
	panic("")
}

func (this *inMemoryDataframe) OuterJoin(with DataFrame, on JoinCondition) DataFrame {
	panic("implement me")
}

func (this *inMemoryDataframe) GetRowCount() int {
	return this.cs.GetRowsCount()
}

func (this *inMemoryDataframe) GetHeaders() []string {
	return this.cs.headers
}

func (this *inMemoryDataframe) Print() {
	tw := tablewriter.NewWriter(os.Stdout)
	tw.SetHeader(this.GetHeaders())

	for i := 0; i < int(math.Min(30, float64(this.GetRowCount()))); i++ {
		r := make([]string, len(this.GetHeaders()))
		row := this.cs.GetRow(i)
		j := 0
		for _, k := range this.GetHeaders() {
			r[j] = fmt.Sprintf("%v", row[k])
			j++
		}
		tw.Append(r)
	}

	tw.Render()
}

func encodeGroupKey(cols []string, row Row) string {
	b := strings.Builder{}
	for _, c := range cols {
		b.WriteString(fmt.Sprintf("%v ", row[c]))
	}
	return b.String()
}

func combineRows(store map[string][]CValue, left Row, right Row, rightMissing bool) {
	dup := make(map[string]bool)

	for k, v := range left {
		dup[k] = true
		store[k] = append(store[k], v)
	}

	for k, v := range right {
		if _, found := dup[k]; found {
			continue
		}
		if rightMissing {
			store[k] = append(store[k], nil)
		} else {
			store[k] = append(store[k], v)
		}
	}
}
