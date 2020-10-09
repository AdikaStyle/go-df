package adapters

import (
	"encoding/csv"
	"github.com/AdikaStyle/go-df/backend"
	"github.com/AdikaStyle/go-df/dataframe"
	"github.com/AdikaStyle/go-df/types"
	"os"
)

func WriteCSV(outPath string, df dataframe.Dataframe, delimiter rune) {
	_ = os.Remove(outPath)
	file, err := os.OpenFile(outPath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	types.PanicOnError(err)

	w := csv.NewWriter(file)
	w.Comma = delimiter

	types.PanicOnError(w.Write(toCSVHeader(df)))

	headers := df.GetHeaders()
	csvRow := make([]string, len(headers))
	df.VisitRows(func(id int, row backend.Row) {
		toCSVRow(headers, row, csvRow)
		types.PanicOnError(w.Write(csvRow))
	})

	w.Flush()
	types.PanicOnError(file.Close())
}

func toCSVHeader(df dataframe.Dataframe) []string {
	var out []string
	for _, h := range df.GetHeaders() {
		out = append(out, h.Name)
	}
	return out
}

func toCSVRow(headers backend.Headers, row backend.Row, csvRow []string) {
	for idx, h := range headers {
		csvRow[idx] = row[h.Name].String()
	}
}
