package go_df

func Not(val Condition) Condition {
	return func(row Row) bool {
		return !val(row)
	}
}

func And(left Condition, right Condition) Condition {
	return func(row Row) bool {
		return left(row) && right(row)
	}
}

func Or(left Condition, right Condition) Condition {
	return func(row Row) bool {
		return left(row) || right(row)
	}
}
