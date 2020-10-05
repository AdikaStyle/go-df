package go_df

import (
	"reflect"
)

var floatType = reflect.TypeOf(float64(0))

func Eq(columnName string, value CValue) Condition {
	return func(row Row) bool {
		return value == row[columnName]
	}
}

func NotEq(columnName string, value CValue) Condition {
	return func(row Row) bool {
		return value != row[columnName]
	}
}

func Gt(columnName string, value CValue) Condition {
	return func(row Row) bool {
		return AsNumber(row[columnName]) > AsNumber(value)
	}
}

func Gte(columnName string, value CValue) Condition {
	return func(row Row) bool {
		return AsNumber(row[columnName]) >= AsNumber(value)
	}
}

func Lt(columnName string, value CValue) Condition {
	return func(row Row) bool {
		return AsNumber(row[columnName]) < AsNumber(value)
	}
}

func Lte(columnName string, value CValue) Condition {
	return func(row Row) bool {
		return AsNumber(row[columnName]) <= AsNumber(value)
	}
}

func In(columnName string, inValues ...CValue) Condition {
	return func(row Row) bool {
		for _, val := range inValues {
			if row[columnName] == val {
				return true
			}
		}
		return false
	}
}

func AsNumber(value CValue) float64 {
	return reflect.ValueOf(value).Convert(floatType).Float()
}
