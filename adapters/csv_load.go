package adapters

import (
	"bytes"
	"encoding/csv"
	"github.com/AdikaStyle/go-df/backend"
	"github.com/AdikaStyle/go-df/dataframe"
	"github.com/AdikaStyle/go-df/types"
	"github.com/palantir/stacktrace"
	"io/ioutil"
)

func ReadCsv(content []byte, delimiter rune) dataframe.Dataframe {
	return ReadCsvWithHeaders(content, delimiter, nil)
}

func ReadCsvWithHeaders(content []byte, delimiter rune, headers backend.Headers) dataframe.Dataframe {
	records, err := readCsv(content, delimiter)
	types.PanicOnError(err)

	if headers == nil {
		headers = AnalyzeDataset(records[0], records[1:])
	}

	be := backend.New(headers)
	populateCSV(be, records)

	return dataframe.New(be)
}

func LoadCSV(path string, delimiter rune) dataframe.Dataframe {
	return LoadCSVWithHeaders(path, delimiter, nil)
}

func LoadCSVWithHeaders(path string, delimiter rune, headers backend.Headers) dataframe.Dataframe {
	content, err := ioutil.ReadFile(path)
	types.PanicOnError(err)

	return ReadCsvWithHeaders(content, delimiter, headers)
}

func populateCSV(be backend.Backend, records [][]string) {
	for i := 1; i < len(records); i++ {
		row := make(backend.Row)
		for hid, h := range be.GetHeaders() {
			row[h.Name] = string2TypedValueKind(records[i][hid], h.Kind)
		}
		be.AppendRow(row)
	}
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
