package dataframe

import (
	"fmt"
	"github.com/AdikaStyle/go-df/backend"
	"github.com/olekukonko/tablewriter"
	"io"
)

type defaultPrintable struct {
	df Dataframe
}

func newDefaultPrintable(df Dataframe) *defaultPrintable {
	return &defaultPrintable{df: df}
}

func (this *defaultPrintable) Print(w io.Writer) Dataframe {
	tw := tablewriter.NewWriter(w)
	headers := this.df.GetHeaders()
	names := onlyNames(headers)

	tw.SetHeader(names)
	tw.SetAutoFormatHeaders(false)

	printRow := make([]string, len(headers))

	this.df.VisitRows(func(id int, row backend.Row) {
		if tw.NumLines() > 30 {
			return
		}

		for idx, h := range names {
			if row[h] != nil {
				printRow[idx] = row[h].String()
			} else {
				printRow[idx] = "?"
			}
		}
		tw.Append(printRow)
	})

	tw.Render()
	fmt.Printf("Dataframe has %d columns and %d rows in total.\n", len(names), this.df.GetRowCount())
	return this.df
}

func onlyNames(headers backend.Headers) []string {
	var headerNames []string
	for _, h := range headers {
		headerNames = append(headerNames, h.Name)
	}
	return headerNames
}
