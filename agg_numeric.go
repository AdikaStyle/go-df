package go_df

import (
	"math"
)

func Sum() Aggregation {
	return func(column []CValue) CValue {
		sum := float64(0)
		for idx := range column {
			sum += AsNumber(column[idx])
		}
		return sum
	}
}

func Avg() Aggregation {
	return func(column []CValue) CValue {
		return Sum()(column).(float64) / float64(len(column))
	}
}

func Min() Aggregation {
	return func(column []CValue) CValue {
		min := math.Inf(1)
		for idx := range column {
			val := AsNumber(column[idx])
			if val < min {
				min = val
			}
		}
		return min
	}
}

func Max() Aggregation {
	return func(column []CValue) CValue {
		max := math.Inf(-1)
		for idx := range column {
			val := AsNumber(column[idx])
			if val > max {
				max = val
			}
		}
		return max
	}
}
