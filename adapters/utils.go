package adapters

import (
	"github.com/AdikaStyle/go-df/backend"
	"github.com/AdikaStyle/go-df/types"
	"github.com/cstockton/go-conv"
	"log"
	"math/rand"
	"time"
)

func AnalyzeDataset(headers []string, data [][]string) backend.Headers {
	tpes := make(map[string]types.TypedValue)
	sample := sampleRows(data, 0.1)

	for _, row := range sample {
		for hid, header := range headers {
			new := parseString(row[hid])
			old := tpes[header]
			if shouldOverrideType(old, new) {
				tpes[header] = new
			}
		}
	}

	var out backend.Headers
	for _, h := range headers {
		out = append(out, backend.Header{
			Name:    h,
			Kind:    tpes[h].Kind(),
			Visible: false,
		})
	}

	return out
}

func shouldOverrideType(old types.TypedValue, new types.TypedValue) bool {
	if old == nil {
		return true
	}

	if new == nil {
		return false
	}

	return new.Precedence() > old.Precedence()
}

func sampleRows(data [][]string, percentage float32) [][]string {
	rowsCount := len(data)
	if rowsCount <= 0 {
		log.Fatalf("canno't sample data from empty array")
	}

	if rowsCount <= 100 {
		return data
	}

	sampleSize := int(percentage * float32(rowsCount))

	out := make([][]string, sampleSize)
	for i := 0; i < sampleSize; i++ {
		sId := rand.Intn(rowsCount)
		out[i] = data[sId]
	}

	return out
}

func parseString(val string) types.TypedValue {
	if isMissing(val) {
		return types.Missing
	}

	/*if res, err := conv.Int64(val); err == nil {
		return types.Integer(res)
	}*/

	if res, err := conv.Float64(val); err == nil {
		return types.Decimal(res)
	}

	if res, err := conv.Bool(val); err == nil {
		return types.Boolean(res)
	}

	if res, err := conv.Time(val); err == nil {
		return types.Datetime(res)
	}

	return types.String(val)
}

func string2TypedValueKind(val string, kind types.TypeKind) types.TypedValue {
	if isMissing(val) {
		return types.Missing
	}

	switch kind {
	case types.KindBoolean:
		var out bool
		types.PanicOnError(conv.Infer(&out, val))
		return types.Boolean(out)
	case types.KindInteger:
		var out int64
		types.PanicOnError(conv.Infer(&out, val))
		return types.Integer(out)
	case types.KindDecimal:
		var out float64
		types.PanicOnError(conv.Infer(&out, val))
		return types.Decimal(out)
	case types.KindDatetime:
		var out time.Time
		err := conv.Infer(&out, val)
		if err != nil {
			return types.Missing
		}
		return types.Datetime(out)
	case types.KindString:
		var out string
		types.PanicOnError(conv.Infer(&out, val))
		return types.String(out)
	case types.KindMissing:
		return types.Missing
	default:
		log.Fatalf("unable to convert string: %v to kind: %s", val, kind)
		return types.Missing

	}
}

func isMissing(val string) bool {
	return val == "null" || val == "" || val == "NULL" || val == "?"
}
