package go_df

import (
	"log"
	"reflect"
)

type Row map[string]CValue

type CValue interface{}

type ColumnStore struct {
	headers []string
	store   map[string][]CValue
	row     Row
}

func NewColumnStore(store map[string][]CValue) *ColumnStore {
	var headers []string
	for k := range store {
		headers = append(headers, k)
	}
	return &ColumnStore{store: store, headers: headers, row: make(map[string]CValue)}
}

func (this *ColumnStore) GetRow(id int) Row {
	for _, header := range this.headers {
		this.row[header] = this.store[header][id]
	}
	return this.row
}

func (this *ColumnStore) RemoveRows(ids ...int) {
	for _, id := range ids {
		for idx := range this.store {
			this.store[idx] = removeCValue(this.store[idx], id)
			println()
		}
	}
}

func (this *ColumnStore) RemoveColumn(name string) {
	nameIdx := -1
	for idx := range this.headers {
		if this.headers[idx] == name {
			nameIdx = idx
		}
	}

	if nameIdx == -1 {
		log.Fatalf("failed to find column named %s while trying to remove column", name)
	}

	removeString(&this.headers, nameIdx)
	delete(this.store, name)
}

func (this *ColumnStore) GetRowsCount() int {
	return len(this.store[this.headers[0]])
}

func removeCValue(a []CValue, i int) []CValue {
	a[i] = a[len(a)-1]
	a[len(a)-1] = reflect.Zero(reflect.TypeOf(a[i])).Interface()
	a = a[:len(a)-1]
	return a
}

func removeString(a *[]string, i int) {
	(*a)[i] = (*a)[len(*a)-1]
	(*a)[len(*a)-1] = ""
	*a = (*a)[:len(*a)-1]
}
