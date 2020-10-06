package go_df

import (
	"bytes"
	"encoding/csv"
	"github.com/cstockton/go-conv"
	"github.com/palantir/stacktrace"
	"io/ioutil"
	"os"
)

func LoadCSV(path string, delimiter rune) (DataFrame, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, stacktrace.RootCause(err)
	}

	records, err := readCsv(content, delimiter)
	if err != nil {
		return nil, stacktrace.Propagate(err, "couldn't read csv file")
	}

	store := make(map[string][]CValue)
	buildStore(records, store)

	return NewInMemoryDataframe(NewColumnStore(store)), nil
}

func WriteCSV(outPath string, delimiter rune, df DataFrame) {
	file, err := os.OpenFile(outPath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	PanicOnError(err)

	w := csv.NewWriter(file)
	w.Comma = delimiter

	headers := df.GetHeaders()
	csvRow := make([]string, len(df.GetHeaders()))

	for idx := range headers {
		csvRow[idx] = headers[idx]
	}
	PanicOnError(w.Write(csvRow))

	df.ForEach(func(row Row) {
		for idx := range headers {
			str, err := conv.String(row[headers[idx]])
			PanicOnError(err)

			csvRow[idx] = str
		}
		PanicOnError(w.Write(csvRow))
	})

	w.Flush()
	PanicOnError(file.Close())
	PanicOnError(w.Error())
}

func buildStore(records [][]string, store map[string][]CValue) {
	headers := records[0]
	for i := 1; i < len(records); i++ {
		for h := 0; h < len(headers); h++ {
			store[headers[h]] = append(store[headers[h]], MatchType(records[i][h]))
		}
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
