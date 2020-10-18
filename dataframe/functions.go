package dataframe

import (
	"github.com/AdikaStyle/go-df/backend"
	"github.com/AdikaStyle/go-df/types"
)

func ReplaceMissingValues(defaultValue types.TypedValue) backend.ApplyFunction {
	return func(value types.TypedValue) types.TypedValue {
		if value == types.Missing {
			return defaultValue
		}
		return value
	}
}
