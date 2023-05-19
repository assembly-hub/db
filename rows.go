package db

type Rows interface {
	ColumnTypes() ([]ColumnType, error)
	Columns() ([]string, error)
	Err() error
	Next() bool
	NextResultSet() bool
	Close() error
	Scan(dest ...any) error
}
