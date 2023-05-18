package db

import "database/sql"

type Rows interface {
	ColumnTypes() ([]*sql.ColumnType, error)
	Columns() ([]string, error)
	Err() error
	Next() bool
	NextResultSet() bool
	Close() error
	Scan(dest ...any) error
}
