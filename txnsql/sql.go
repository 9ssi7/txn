package txnsql

import (
	"context"
	"database/sql"

	"github.com/9ssi7/txn"
)

type SqlDbTx interface {
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	Prepare(query string) (*sql.Stmt, error)
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	Exec(query string, args ...any) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	Query(query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	QueryRow(query string, args ...any) *sql.Row
}

// SqlAdapter is the interface for interacting with SQL databases within a transaction.
// It extends the txn.Adapter interface to provide additional SQL-specific functionality.
type SqlAdapter interface {
	txn.Adapter

	// Returns current transaction if it exists.
	GetCurrent() SqlDbTx
}

// New creates a new SqlAdapter instance using the provided *sql.DB.
func New(db *sql.DB) SqlAdapter {
	return &sqlAdapter{db: db}
}

type sqlAdapter struct {
	db *sql.DB
	tx *sql.Tx
}

func (a *sqlAdapter) Begin(ctx context.Context) error {
	tx, err := a.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	a.tx = tx
	return nil
}

func (a *sqlAdapter) Commit(_ context.Context) error {
	if a.tx == nil {
		return nil
	}
	err := a.tx.Commit()
	a.tx = nil
	return err
}

func (a *sqlAdapter) Rollback(_ context.Context) error {
	if a.tx == nil {
		return nil
	}
	err := a.tx.Rollback()
	a.tx = nil
	return err
}

func (a *sqlAdapter) End(_ context.Context) {
	if a.tx != nil {
		a.tx = nil
	}
}

func (a *sqlAdapter) GetCurrent() SqlDbTx {
	if a.tx == nil {
		return a.db
	}
	return a.tx
}
