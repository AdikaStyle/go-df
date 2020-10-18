package adapters

import (
	"database/sql"
	"github.com/AdikaStyle/go-df/backend"
	"github.com/AdikaStyle/go-df/dataframe"
	"github.com/AdikaStyle/go-df/types"
)

func LoadSQL(rows *sql.Rows) dataframe.Dataframe {
	if rows == nil {
		panic("unable to parse nil rows (LoadSQL)")
	}
	defer rows.Close()

	if err := rows.Err(); err != nil {
		panic(err)
	}

	cols, err := rows.Columns()
	if err != nil {
		panic(err)
	}

	csvRows := parseRows(cols, rows)

	headers := AnalyzeDataset(cols, csvRows)
	be := backend.New(headers)
	df := dataframe.New(be)

	row := make(backend.Row)

	for _, r := range csvRows {
		for idx, col := range cols {
			row[col] = string2TypedValueKind(r[idx], headers[idx].Kind)
			if be.GetHeaders()[idx].Kind == types.KindUnknown {
				be.GetHeaders()[idx].Kind = row[col].Kind()
			}
		}
		be.AppendRow(row)
	}

	return df
}

func parseRows(headers []string, rows *sql.Rows) [][]string {
	rawResult := make([][]byte, len(headers))
	result := make([]string, len(headers))

	dest := make([]interface{}, len(headers))
	for i := range rawResult {
		dest[i] = &rawResult[i]
	}

	var out [][]string
	for rows.Next() {
		outRow := make([]string, len(headers))
		err := rows.Scan(dest...)
		if err != nil {
			panic(err)
		}

		for i, raw := range rawResult {
			if raw == nil {
				result[i] = "\\N"
			} else {
				result[i] = string(raw)
			}
		}

		for idx := range result {
			outRow[idx] = result[idx]
		}

		out = append(out, outRow)
	}

	return out
}
