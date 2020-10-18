package adapters

import (
	"database/sql"
	"fmt"
	"github.com/AdikaStyle/go-df/backend"
	"github.com/AdikaStyle/go-df/dataframe"
	"github.com/AdikaStyle/go-df/types"
	"strings"
)

func WriteSQL(tableName string, df dataframe.Dataframe, db *sql.DB) {
	tx, err := db.Begin()
	types.PanicOnError(err)

	stmt, err := tx.Prepare(buildSqlInsertStmt(tableName, df))
	types.PanicOnError(err)

	df.VisitRows(func(id int, row backend.Row) {
		_, err := stmt.Exec(sqlValues(df, row)...)
		types.PanicOnError(err)
	})

	types.PanicOnError(tx.Commit())
}

func buildSqlInsertStmt(tableName string, df dataframe.Dataframe) string {
	return fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, sqlHeaders(df), sqlArgs(df))
}

func sqlHeaders(df dataframe.Dataframe) string {
	headers := make([]string, len(df.GetHeaders()))
	for i, h := range df.GetHeaders() {
		headers[i] = h.Name
	}

	return strings.Join(headers, ",")
}

func sqlArgs(df dataframe.Dataframe) string {
	headers := make([]string, len(df.GetHeaders()))
	for i := range headers {
		headers[i] = "?"
	}

	return strings.Join(headers, ",")
}

func sqlValues(df dataframe.Dataframe, row backend.Row) []interface{} {
	values := make([]interface{}, len(row))
	for i, h := range df.GetHeaders() {
		val := row[h.Name]
		values[i] = val.NativeType()
	}
	return values
}
