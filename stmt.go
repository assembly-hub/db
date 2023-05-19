package db

import (
	"context"
)

type Stmt interface {
	Close() error
	ExecContext(ctx context.Context, args ...any) (Result, error)
	QueryContext(ctx context.Context, args ...any) (Rows, error)
	QueryRowContext(ctx context.Context, args ...any) Row
}
