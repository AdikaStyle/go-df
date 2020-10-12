package adapters

import (
	"database/sql"
	"github.com/AdikaStyle/go-df/backend"
	"github.com/AdikaStyle/go-df/dataframe"
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

	headers := buildHeaders(cols)
	be := backend.New(headers)
	df := dataframe.New(be)

	row := make(backend.Row)
	iterateRows(cols, rows, func(in []string) {
		for idx, col := range cols {
			row[col] = parseString(in[idx])
		}
		be.AppendRow(row)
	})

	return df
}

func iterateRows(cols []string, rows *sql.Rows, onRow func(in []string)) {
	rawResult := make([][]byte, len(cols))
	result := make([]string, len(cols))

	dest := make([]interface{}, len(cols))
	for i := range rawResult {
		dest[i] = &rawResult[i]
	}

	for rows.Next() {
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

		onRow(result)
	}
}
