package manticore

const (
	ColumnText      ColumnType = "text"    // alias: string
	ColumnInteger   ColumnType = "integer" // alias: int, uint
	ColumnBitInt    ColumnType = "bigint"
	ColumnFloat     ColumnType = "float"
	ColumnMulti     ColumnType = "multi"
	ColumnMulti64   ColumnType = "multi64"
	ColumnBool      ColumnType = "bool"
	ColumnJSON      ColumnType = "json"
	ColumnString    ColumnType = "string"
	ColumnTimestamp ColumnType = "timestamp"
)

type ColumnType string

type Column struct {
	Field      string
	Type       ColumnType
	Properties string
}
