package adapters

import (
	"github.com/AdikaStyle/go-df/backend"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestLoadStructs(t *testing.T) {
	structs := getStructs()
	df := LoadStructs(structs, "df")

	assert.NotNil(t, df)
	assert.EqualValues(t, 5, len(df.GetHeaders()))
	assert.EqualValues(t, "name", df.GetHeaders()[0].Name)
	assert.EqualValues(t, "age", df.GetHeaders()[1].Name)
	assert.EqualValues(t, "salary", df.GetHeaders()[2].Name)
	assert.EqualValues(t, "join_date", df.GetHeaders()[3].Name)
	assert.EqualValues(t, "is_active", df.GetHeaders()[4].Name)

	assert.EqualValues(t, 2, df.GetRowCount())
	df.VisitRows(func(id int, row backend.Row) {
		expected := structs[id]
		assert.EqualValues(t, expected.Name, row["name"].NativeType())
		assert.EqualValues(t, expected.Age, row["age"].NativeType())
		assert.EqualValues(t, expected.Salary, row["salary"].NativeType())
		assert.EqualValues(t, expected.JoinDate, row["join_date"].NativeType())
		assert.EqualValues(t, expected.IsActive, row["is_active"].NativeType())
	})
}

func getStructs() []demoStruct {
	return []demoStruct{
		{
			Name:     "User1",
			Age:      10,
			Salary:   1203.3,
			JoinDate: time.Date(2020, time.May, 28, 20, 00, 00, 00, time.UTC),
			IsActive: true,
		},
		{
			Name:     "User2",
			Age:      20,
			Salary:   1203.311,
			JoinDate: time.Date(2020, time.May, 28, 20, 00, 00, 00, time.UTC),
			IsActive: false,
		},
	}
}

type demoStruct struct {
	Name     string    `df:"name"`
	Age      int       `df:"age"`
	Salary   float32   `df:"salary"`
	JoinDate time.Time `df:"join_date"`
	IsActive bool      `df:"is_active"`
	Ignore   int       `df:"-"`
}
