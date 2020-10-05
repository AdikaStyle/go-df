package go_df

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInMemoryDataframe_Select(t *testing.T) {
	var exampleDataset = map[string][]CValue{
		"name": {"Abraham", "Yitzack", "Yaakov"},
		"age":  {102, 203, 99},
		"brit": {true, true, true},
	}

	df := NewInMemoryDataframe(NewColumnStore(exampleDataset))
	assert.EqualValues(t, len(df.cs.store[df.cs.headers[0]]), len(exampleDataset["name"]))

	df.Select("name", "brit")

	df.Print()

	assert.Contains(t, df.cs.headers, "name")
	assert.Contains(t, df.cs.headers, "brit")
	assert.EqualValues(t, len(df.cs.store[df.cs.headers[0]]), len(exampleDataset["name"]))
}

func TestInMemoryDataframe_Filter(t *testing.T) {
	var exampleDataset = map[string][]CValue{
		"name": {"Abraham", "Yitzack", "Yaakov"},
		"age":  {102, 203, 99},
		"brit": {true, true, true},
	}

	df := NewInMemoryDataframe(NewColumnStore(exampleDataset))

	df.Filter(func(row Row) bool {
		return row["age"].(int) > 100
	})

	df.Print()

	assert.Contains(t, df.cs.headers, "name")
}

func TestInMemoryDataframe_Concat(t *testing.T) {
	var exampleDataset1 = map[string][]CValue{
		"name": {"Abraham", "Yaakov"},
		"age":  {102, 99},
		"brit": {true, true},
	}

	var exampleDataset2 = map[string][]CValue{
		"name": {"Yitzack"},
		"age":  {203},
		"brit": {true},
	}

	df1 := NewInMemoryDataframe(NewColumnStore(exampleDataset1))
	df2 := NewInMemoryDataframe(NewColumnStore(exampleDataset2))

	df1.Concat(df2)

	df1.Print()

	assert.EqualValues(t, 3, df1.cs.GetRowsCount())
}

func TestInMemoryDataframe_Split(t *testing.T) {
	var exampleDataset = map[string][]CValue{
		"name": {"Abraham", "Yitzack", "Yaakov"},
		"age":  {102, 203, 99},
		"brit": {true, true, true},
	}

	df := NewInMemoryDataframe(NewColumnStore(exampleDataset))

	onTrue, onFalse := df.Split(func(row Row) bool {
		return row["age"].(int) > 100
	})

	onTrue.Print()
	onFalse.Print()

	assert.EqualValues(t, 2, onTrue.GetRowCount())
	assert.EqualValues(t, 1, onFalse.GetRowCount())
}

func TestInMemoryDataframe_GroupBy(t *testing.T) {
	var exampleDataset = map[string][]CValue{
		"sku":      {"A0", "A0", "A0", "A1", "A1", "A1", "A1", "A1"},
		"zone":     {"IL", "IL", "US", "IL", "IL", "EU", "EU", "EU"},
		"quantity": {2, 1, 1, 5, 5, 2, 1, 1},
		"sales":    {200, 100, 100, 100, 800, 300, 200, 100},
	}

	df := NewInMemoryDataframe(NewColumnStore(exampleDataset))
	gdf := df.GroupBy(Cols{"sku", "zone"}, Aggs{"quantity": Sum(), "sales": Sum()})

	row := gdf.(*inMemoryDataframe).cs.GetRow(0)

	gdf.Print()
	assert.NotNil(t, row)
}

func TestInMemoryDataframe_InnerJoin(t *testing.T) {
	var exampleDatasetFact = map[string][]CValue{
		"sku":      {"A0", "A0", "A0", "A1", "A1", "A1", "A1", "A1"},
		"zone":     {"IL", "IL", "US", "IL", "IL", "EU", "EU", "EU"},
		"quantity": {2, 1, 1, 5, 5, 2, 1, 1},
		"sales":    {200, 100, 100, 100, 800, 300, 200, 100},
	}

	var exampleDatasetDim = map[string][]CValue{
		"sku":    {"A0", "A1", "A2"},
		"name":   {"Car", "Airplane", "Spaceship"},
		"engine": {"V7", "Jet", "Neuron"},
	}

	dfFact := NewInMemoryDataframe(NewColumnStore(exampleDatasetFact))
	dfDim := NewInMemoryDataframe(NewColumnStore(exampleDatasetDim))

	joinedDf := dfFact.InnerJoin(dfDim, func(left Row, right Row) bool {
		return left["sku"] == right["sku"]
	})

	joinedDf.Print()
	assert.NotNil(t, joinedDf)
}

func TestInMemoryDataframe_LeftJoin(t *testing.T) {
	var exampleDatasetFact = map[string][]CValue{
		"sku":      {"A0", "A0", "A0", "A1", "A1", "A1", "A1", "A1"},
		"zone":     {"IL", "IL", "US", "IL", "IL", "EU", "EU", "EU"},
		"quantity": {2, 1, 1, 5, 5, 2, 1, 1},
		"sales":    {200, 100, 100, 100, 800, 300, 200, 100},
	}

	var exampleDatasetDim = map[string][]CValue{
		"sku":    {"A0", "A2"},
		"name":   {"Car", "Spaceship"},
		"engine": {"V7", "Neuron"},
	}

	dfFact := NewInMemoryDataframe(NewColumnStore(exampleDatasetFact))
	dfDim := NewInMemoryDataframe(NewColumnStore(exampleDatasetDim))

	joinedDf := dfFact.LeftJoin(dfDim, func(left Row, right Row) bool {
		return left["sku"] == right["sku"]
	})
	joinedDf.Print()
	assert.NotNil(t, joinedDf)
}

func TestInMemoryDataframe_RightJoin(t *testing.T) {
	var exampleDatasetFact = map[string][]CValue{
		"sku":      {"A0", "A0", "A0", "A1", "A1", "A1", "A1", "A1"},
		"zone":     {"IL", "IL", "US", "IL", "IL", "EU", "EU", "EU"},
		"quantity": {2, 1, 1, 5, 5, 2, 1, 1},
		"sales":    {200, 100, 100, 100, 800, 300, 200, 100},
	}

	var exampleDatasetDim = map[string][]CValue{
		"sku":    {"A0", "A2"},
		"name":   {"Car", "Spaceship"},
		"engine": {"V7", "Neuron"},
	}

	dfFact := NewInMemoryDataframe(NewColumnStore(exampleDatasetFact))
	dfDim := NewInMemoryDataframe(NewColumnStore(exampleDatasetDim))

	joinedDf := dfDim.RightJoin(dfFact, func(left Row, right Row) bool {
		return left["sku"] == right["sku"]
	})
	joinedDf.Print()
	assert.NotNil(t, joinedDf)
}
