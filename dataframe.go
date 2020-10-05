package go_df

type Condition func(row Row) bool

type JoinCondition func(left Row, right Row) bool

type Aggregation func(column []CValue) CValue

type Aggs map[string]Aggregation
type Cols []string

type DataFrame interface {
	Select(columns ...string) DataFrame
	Filter(cond Condition) DataFrame
	Concat(with DataFrame) DataFrame
	LeftJoin(with DataFrame, on JoinCondition) DataFrame
	RightJoin(with DataFrame, on JoinCondition) DataFrame
	InnerJoin(with DataFrame, on JoinCondition) DataFrame
	OuterJoin(with DataFrame, on JoinCondition) DataFrame
	Split(cond Condition) (onTrue DataFrame, onFalse DataFrame)
	GroupBy(columns []string, aggregations map[string]Aggregation) DataFrame

	GetRowCount() int
	GetHeaders() []string
	Print()
}
