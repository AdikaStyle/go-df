package adapters

import (
	"bytes"
	"encoding/csv"
	"github.com/AdikaStyle/go-df/backend"
	"github.com/AdikaStyle/go-df/dataframe"
	"github.com/AdikaStyle/go-df/types"
	"github.com/cstockton/go-conv"
	"github.com/palantir/stacktrace"
	"io/ioutil"
)

func LoadCSV(path string, delimiter rune) dataframe.Dataframe {
	content, err := ioutil.ReadFile(path)
	types.PanicOnError(err)

	records, err := readCsv(content, delimiter)
	types.PanicOnError(err)

	headers := buildHeaders(records)

	be := backend.New(headers)
	populateCSV(be, records)

	return dataframe.New(be)
}

func populateCSV(be backend.Backend, records [][]string) {
	for i := 1; i < len(records); i++ {
		row := make(backend.Row)
		for hid, h := range be.GetHeaders() {
			row[h.Name] = parseString(records[i][hid])
		}
		be.AppendRow(row)
	}
}

func buildHeaders(csvResult [][]string) backend.Headers {
	var headers backend.Headers
	for _, h := range csvResult[0] {
		headers = append(headers, backend.Header{
			Name:    h,
			Kind:    types.KindUnknown,
			Visible: false,
		})
	}

	for idx, v := range csvResult[1] {
		headers[idx].Kind = parseString(v).Kind()
	}

	return headers
}

func parseString(val string) types.TypedValue {
	if isMissing(val) {
		return types.Missing
	}

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

func readCsv(content []byte, delimiter rune) ([][]string, error) {
	r := csv.NewReader(bytes.NewBuffer(content))
	r.Comma = delimiter
	records, err := r.ReadAll()
	if err != nil {
		return nil, stacktrace.RootCause(err)
	}
	return records, nil
}

func isMissing(val string) bool {
	return val == "null" || val == "" || val == "NULL" || val == "?"
}
