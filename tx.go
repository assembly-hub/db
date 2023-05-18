package db

import "context"

type Tx interface {
	BaseExecutor

	Commit() error
	Rollback() error
	StmtContext(ctx context.Context, stmt Stmt) Stmt
}
