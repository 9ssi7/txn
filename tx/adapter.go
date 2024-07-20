package tx

import "context"

// Adapter defines the interface for interacting with a specific data source
// within a transaction.
type Adapter interface {

	// Begin starts a transaction on the data source.
	Begin(ctx context.Context) error

	// Commit commits the transaction on the data source.
	Commit(ctx context.Context) error

	// Rollback rolls back the transaction on the data source.
	Rollback(ctx context.Context) error

	// End is called at the end of a transaction to clean up any resources.
	// It's called regardless of whether the transaction was committed or rolled back.
	End(ctx context.Context)
}
