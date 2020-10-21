package adapters

import (
	"database/sql"
	"fmt"
	"github.com/AdikaStyle/go-df/backend"
	"github.com/AdikaStyle/go-df/dataframe"
	"github.com/AdikaStyle/go-df/types"
	"log"
	"strings"
)

func WriteSQL(tableName string, df dataframe.Dataframe, db *sql.DB, batchSize int) {
	sb := newSqlBatcher(tableName, df, db, batchSize)
	sb.Do()
}

type sqlBatcher struct {
	db          *sql.DB
	insertQuery string
	batchSize   int
	df          dataframe.Dataframe

	stmt *sql.Stmt
	tx   *sql.Tx
}

func newSqlBatcher(tableName string, df dataframe.Dataframe, db *sql.DB, batchSize int) *sqlBatcher {
	query := buildSqlInsertStmt(tableName, df)
	return &sqlBatcher{db: db, batchSize: batchSize, df: df, insertQuery: query}
}

func (sb *sqlBatcher) Do() {
	var idx = 0
	var batchIdx = 1
	total := sb.df.GetRowCount()
	sb.nextStmt()
	sb.df.VisitRows(func(id int, row backend.Row) {
		idx++

		_, err := sb.stmt.Exec(sqlValues(sb.df, row)...)
		types.PanicOnError(err)

		if idx%sb.batchSize == 0 {
			log.Printf("Writing batch num: %d, %d/%d\n", batchIdx, batchIdx*sb.batchSize, total)
			sb.flush()
			sb.nextStmt()
			idx = 0
			batchIdx++
		}
	})

	if idx > 0 {
		batchIdx++
		log.Printf("Writing batch num: %d, %d/%d\n", batchIdx, total-(batchIdx-1)*sb.batchSize, total)
		sb.flush()
		idx = 0
	}
}

func (sb *sqlBatcher) nextStmt() {
	tx, err := sb.db.Begin()
	types.PanicOnError(err)

	stmt, err := tx.Prepare(sb.insertQuery)
	types.PanicOnError(err)

	sb.tx = tx
	sb.stmt = stmt
}

func (sb *sqlBatcher) flush() {
	types.PanicOnError(sb.tx.Commit())
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
