package db

import (
	"context"
	"database/sql"
)

type BaseExecutor interface {
	PrepareContext(ctx context.Context, query string) (Stmt, error)

	ExecContext(ctx context.Context, query string, args ...any) (Result, error)

	QueryContext(ctx context.Context, query string, args ...any) (Rows, error)

	QueryRowContext(ctx context.Context, query string, args ...any) Row
}

type Executor interface {
	BaseExecutor

	BeginTx(ctx context.Context, opts *sql.TxOptions) (Tx, error)
}
