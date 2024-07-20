package tx

import "context"

// Tx is the core interface for managing transactions.
type Tx interface {

	// Begin starts a new transaction.
	Begin(ctx context.Context) error

	// Commit commits the transaction.
	Commit(ctx context.Context) error

	// Rollback rolls back the transaction.
	Rollback(ctx context.Context) error

	// Cancel cancels the transaction. This is useful in cases where you want
	// to abort the transaction without waiting for the context to be canceled.
	Cancel(ctx context.Context)

	// Register registers an adapter for a specific data source to participate
	// in the transaction.
	Register(Adapter)
}

// New creates a new Tx instance.
func New() Tx {
	return &txn{}
}

type txn struct {
	adapters []Adapter
}

func (t *txn) Register(a Adapter) {
	t.adapters = append(t.adapters, a)
}

func (t *txn) Cancel(ctx context.Context) {
	for _, a := range t.adapters {
		a.Rollback(ctx)
	}
}

func (t *txn) Begin(ctx context.Context) error {
	for _, a := range t.adapters {
		if err := a.Begin(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (t *txn) Commit(ctx context.Context) error {
	for _, a := range t.adapters {
		if err := a.Commit(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (t *txn) Rollback(ctx context.Context) error {
	for _, a := range t.adapters {
		if err := a.Rollback(ctx); err != nil {
			return err
		}
	}
	return nil
}
