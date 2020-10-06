package go_df

import (
	"github.com/cstockton/go-conv"
)

type CValue interface{}

func MatchType(in string) CValue {
	if res, err := conv.Float64(in); err == nil {
		return res
	} else if res, err := conv.Int64(in); err == nil {
		return res
	} else if res, err := conv.Bool(in); err == nil {
		return res
	} else if res, err := conv.Time(in); err == nil {
		return res
	}

	if in == "NULL" || in == "null" || in == "" {
		return nil
	}

	return in
}
