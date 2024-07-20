package txnredis

import (
	"context"

	"github.com/9ssi7/txn"
	"github.com/redis/go-redis/v9"
)

type redisAdapter struct {
	db   *redis.Client
	pipe redis.Pipeliner
}

func New(db *redis.Client) txn.Adapter {
	return &redisAdapter{db: db}
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
}
