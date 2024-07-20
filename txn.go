package txn

import "context"

type Tx interface {
	Begin(ctx context.Context) error
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
	Cancel(ctx context.Context)

	Register(Adapter)
}

type txn struct {
	adapters []Adapter
}

func New() Tx {
	return &txn{}
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
