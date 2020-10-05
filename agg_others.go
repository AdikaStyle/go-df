package go_df

func Count() Aggregation {
	return func(column []CValue) CValue {
		return len(column)
	}
}

func CountDistinct() Aggregation {
	return func(column []CValue) CValue {
		dup := make(map[CValue]bool)
		for idx := range column {
			dup[column[idx]] = true
		}
		return len(dup)
	}
}

func First() Aggregation {
	return func(column []CValue) CValue {
		return column[0]
	}
}

func Last() Aggregation {
	return func(column []CValue) CValue {
		return column[len(column)-1]
	}
}
