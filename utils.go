package go_df

func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
