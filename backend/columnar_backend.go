package backend

import (
	"github.com/cstockton/go-conv"
	"log"
	"math/rand"
)

type columnarBackend struct {
	headers Headers
	columns Columns
	row     Row

	filtered int
}

func NewColumnarBackend(headers Headers) *columnarBackend {
	columns := make(Columns)
	for _, h := range headers {
		columns[h.Name] = nil
	}

	return &columnarBackend{
		columns:  columns,
		headers:  headers,
		row:      make(Row),
		filtered: 0,
	}
}

func (this *columnarBackend) GetHeaders() Headers {
	return this.headers
}

func (this *columnarBackend) GetRowCount() int {
	return len(this.columns[this.headers[0].Name]) - this.filtered
}

func (this *columnarBackend) GetRow(id int) Row {
	for _, header := range this.headers {
		this.row[header.Name] = this.columns[header.Name][id].TypedValue
	}
	return this.row
}

func (this *columnarBackend) AppendRow(row ...Row) {
	for _, r := range row {
		for _, h := range this.headers {
			this.columns[h.Name] = append(this.columns[h.Name], Cell{
				TypedValue: r[h.Name],
				filtered:   false,
			})
		}
	}
}

func (this *columnarBackend) ForEachRow(visitor RowVisitor) {
	for i := 0; i < this.getPhysicalCount(); i++ {
		if !this.columns[this.headers[0].Name][i].filtered {
			visitor(i, this.GetRow(i))
		}
	}
}

func (this *columnarBackend) ForEachColumn(visitor ColumnVisitor) {
	for _, header := range this.headers {
		visitor(header.Name, this.columns[header.Name])
	}
}

func (this *columnarBackend) RemoveRows(ids ...int) {
	for _, id := range ids {
		for _, header := range this.headers {
			this.columns[header.Name][id].filtered = true
		}
		this.filtered += 1
	}
}

func (this *columnarBackend) RemoveColumn(name string) {
	nameIdx := -1
	for idx := range this.headers {
		if this.headers[idx].Name == name {
			nameIdx = idx
		}
	}

	if nameIdx == -1 {
		log.Fatalf("failed to find column named %s while trying to remove column", name)
	}

	removeString(&this.headers, nameIdx)
	delete(this.columns, name)
}

func (this *columnarBackend) RenameColumn(old string, new string) {
	nameIdx := -1
	for idx := range this.headers {
		if this.headers[idx].Name == old {
			nameIdx = idx
		}
	}

	if nameIdx == -1 {
		log.Fatalf("failed to find column named %s while trying to rename column", old)
	}

	this.headers[nameIdx].Name = new

	this.columns[new] = this.columns[old]
	delete(this.columns, old)
}

func (this *columnarBackend) SortByColumn(column string, order Ordering) {
	sortCol := this.columns[column]

	this.quickSort(sortCol[:], 0, bool(order), func(idx1, idx2 int) {
		for _, k := range this.headers {
			this.columns[k.Name][idx1], this.columns[k.Name][idx2] =
				this.columns[k.Name][idx2], this.columns[k.Name][idx1]
		}
	})

	return
}

func (this *columnarBackend) ConstructNew(headers Headers) Backend {
	return NewColumnarBackend(headers)
}

func (this *columnarBackend) SetColumns(columns Columns) {
	this.columns = columns
	this.headers = columns.GetHeaders()
	this.filtered = 0
}

func (this *columnarBackend) quickSort(a Column, offset int, order bool, fn swapFunc) {
	if len(a) < 2 {
		return
	}

	left, right := 0, len(a)-1
	pivot := rand.Int() % len(a)

	fn(pivot+offset, right+offset)

	for i := range a {
		ai, err := conv.Int64(a[i].String())
		if err != nil {
			panic(err)
		}

		aright, err := conv.Int64(a[right].String())
		if err != nil {
			panic(err)
		}

		if order {
			if ai < aright {
				fn(left+offset, i+offset)
				left++
			}
		} else {
			if ai > aright {
				fn(left+offset, i+offset)
				left++
			}
		}
	}

	fn(left+offset, right+offset)

	this.quickSort(a[:left], offset, order, fn)
	this.quickSort(a[left+1:], offset+left+1, order, fn)
	return
}

type swapFunc = func(idx1, idx2 int)

func (this *columnarBackend) getPhysicalCount() int {
	return len(this.columns[this.headers[0].Name])
}

func removeString(a *Headers, i int) {
	(*a)[i] = (*a)[len(*a)-1]
	(*a)[len(*a)-1] = Header{}
	*a = (*a)[:len(*a)-1]
}
