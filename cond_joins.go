package go_df

func On(column string, columns ...string) JoinCondition {
	return func(left Row, right Row) bool {
		var out = left[column] == right[column]
		for _, col := range columns {
			out = out && (left[col] == right[col])
		}
		return out
	}
}
