package txn

import (
	"context"
)

type Repo[DB any] interface {
	WithTxn(d DB)
	GetCurrentDB() DB
	ClearTxn()
}

// Callback is a function type that accepts a database connection and returns an error.
// It is used to add a transactional operation to a transaction.
// The operation will be executed when the transaction is committed.
type Callback[DB any] func(DB) error

// Txn is an interface that represents a transaction.
// It provides methods to add a transactional operation and to commit the transaction.
type Txn[DB any] interface {

	// Add adds a transactional operation to the transaction.
	// The operation will be executed when the transaction is committed.
	// The operation is represented by a callback function that accepts a database connection and returns an error.
	// The callback function should perform the transactional operation.
	Add(cb Callback[DB])

	// Transaction commits the transaction.
	// It executes all the transactional operations that were added to the transaction.
	// If any of the operations return an error, the transaction is rolled back.
	Transaction(ctx context.Context) error
}
