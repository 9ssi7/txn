package txnredis

import (
	"context"
	"errors"
	"testing"

	"github.com/go-redis/redismock/v9"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert" // Or use another assertion library
)

func TestNew(t *testing.T) {
	client := redis.NewClient(&redis.Options{})
	adapter := New(client)
	assert.NotNil(t, adapter)
	assert.IsType(t, &redisAdapter{}, adapter)
}

func TestBegin(t *testing.T) {
	db, _ := redismock.NewClientMock()
	adapter := &redisAdapter{db: db}

	err := adapter.Begin(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, adapter.pipe)
}

func TestCommit_Success(t *testing.T) {
	db, mock := redismock.NewClientMock()
	adapter := &redisAdapter{db: db, pipe: db.TxPipeline()}

	mock.ExpectTxPipelineExec().SetVal([]interface{}{"OK"})
	err := adapter.Commit(context.Background())
	assert.Nil(t, err)
}

func TestCommit_Failure(t *testing.T) {
	db, mock := redismock.NewClientMock()
	adapter := &redisAdapter{db: db, pipe: db.TxPipeline()}

	mock.ExpectTxPipelineExec().SetErr(errors.New("Redis error"))
	adapter.pipe.Set(context.Background(), "key", "value", 0)
	err := adapter.Commit(context.Background())
	assert.Error(t, err)
}

func TestRollback(t *testing.T) {
	db, _ := redismock.NewClientMock()
	adapter := &redisAdapter{db: db, pipe: db.TxPipeline()}
	adapter.pipe.Set(context.Background(), "key", "value", 0) // Add a command to the pipeline

	err := adapter.Rollback(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, 0, adapter.pipe.Len()) // Pipeline should be cleared after rollback
}

func TestEnd(t *testing.T) {
	db, _ := redismock.NewClientMock()
	adapter := &redisAdapter{db: db, pipe: db.TxPipeline()}
	adapter.pipe.Set(context.Background(), "key", "value", 0) // Add a command to the pipeline

	adapter.End(context.Background())      // Discard the pipeline
	assert.Equal(t, 0, adapter.pipe.Len()) // Pipeline should be empty
}
