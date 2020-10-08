package aggs

import (
	"github.com/AdikaStyle/go-df/backend"
	"github.com/AdikaStyle/go-df/types"
	"math"
)

func Sum() Aggregation {
	return func(column backend.Column) types.TypedValue {
		sum := float64(0)
		var value float64
		for idx := range column {
			column[idx].Cast(&value)
			sum += value
		}
		return types.Decimal(sum)
	}
}

func Avg() Aggregation {
	return func(column backend.Column) types.TypedValue {
		return Sum()(column).(types.Decimal) / types.Decimal(len(column))
	}
}

func Min() Aggregation {
	return func(column backend.Column) types.TypedValue {
		min := math.Inf(1)
		for idx := range column {
			var val float64
			(column[idx]).Cast(&val)
			if val < min {
				min = val
			}
		}
		return types.Decimal(min)
	}
}

func Max() Aggregation {
	return func(column backend.Column) types.TypedValue {
		max := math.Inf(-1)
		for idx := range column {
			var val float64
			(column[idx]).Cast(&val)
			if val > max {
				max = val
			}
		}
		return types.Decimal(max)
	}
}
