package txn

import "context"

type Adapter interface {
	Begin(ctx context.Context) error
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
	End(ctx context.Context)
}
