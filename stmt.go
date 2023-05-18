package db

import (
	"context"
	"database/sql"
)

type Stmt interface {
	Close() error
	ExecContext(ctx context.Context, args ...any) (sql.Result, error)
	QueryContext(ctx context.Context, args ...any) (Rows, error)
	QueryRowContext(ctx context.Context, args ...any) Row
}
