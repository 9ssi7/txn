package txnredis

import (
	"context"

	"github.com/9ssi7/txn"
	"github.com/redis/go-redis/v9"
)

// RAdapter is the interface for interacting with Redis within a transaction.
// It extends the txn.Adapter interface to provide additional Redis-specific functionality.
type RAdapter interface {
	txn.Adapter

	// GetCurrent returns the current redis.Cmdable instance to use for executing Redis commands.
	// Depending on the transaction state, this may be the underlying redis.Client or a redis.Pipeliner.
	GetCurrent(ctx context.Context) redis.Cmdable
}

// New creates a new RAdapter instance using the provided redis.Client.
func New(db *redis.Client) RAdapter {
	return &redisAdapter{db: db}
}

type redisAdapter struct {
	db   *redis.Client
	pipe redis.Pipeliner
}

func (a *redisAdapter) Begin(ctx context.Context) error {
	a.pipe = a.db.TxPipeline()
	return nil
}

func (a *redisAdapter) Commit(ctx context.Context) error {
	_, err := a.pipe.Exec(ctx)
	return err
}

func (a *redisAdapter) Rollback(ctx context.Context) error {
	a.pipe.Discard()
	return nil
}

func (a *redisAdapter) End(ctx context.Context) {
	a.pipe.Discard()
	a.pipe = nil
}

func (a *redisAdapter) GetCurrent(ctx context.Context) redis.Cmdable {
	if a.pipe == nil {
		return a.db
	}
	return a.pipe
}
