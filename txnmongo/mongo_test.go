package txnmongo

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

func TestNew(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	adapter := New(mt.Client)
	assert.NotNil(t, adapter)
	assert.IsType(t, &mongoAdapter{}, adapter)
}

func TestBegin_Success(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock).ShareClient(true))

	// Mock successful session start and transaction start
	mt.AddMockResponses(mtest.CreateSuccessResponse(), mtest.CreateSuccessResponse())

	b := true
	ses, _ := mt.Client.StartSession(&options.SessionOptions{
		Snapshot:          &b,
		CausalConsistency: &b,
	})
	adapter := &mongoAdapter{client: mt.Client, sess: ses}
	err := adapter.Begin(context.Background())

	assert.Nil(t, err)
	assert.NotNil(t, adapter.sess)
}

func TestBegin_SessionStartError(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock).ShareClient(true))

	// Mock session start failure
	mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{}))
	b := true
	adapter := mongoAdapter{client: mt.Client, sesOption: &options.SessionOptions{
		Snapshot:          &b,
		CausalConsistency: &b,
	}}
	err := adapter.Begin(context.Background())
	assert.Error(t, err)
}

func TestBegin_TransactionStartError(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock).ShareClient(true))

	// Mock transaction start failure
	mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{}))

	adapter := mongoAdapter{client: mt.Client, txOption: &options.TransactionOptions{
		WriteConcern: writeconcern.Unacknowledged(),
	}}
	err := adapter.Begin(context.Background())
	assert.Error(t, err)
}

func TestCommit_Success(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock).ShareClient(true))
	mt.AddMockResponses(mtest.CreateSuccessResponse())
	ses, _ := mt.Client.StartSession()
	adapter := &mongoAdapter{
		client: mt.Client,
		sess:   ses,
	}
	adapter.Begin(context.Background()) // Start transaction

	err := adapter.Commit(context.Background())
	assert.Nil(t, err)
}

func TestCommit_NoSession(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock).ShareClient(true))

	adapter := &mongoAdapter{client: mt.Client} // No session set
	err := adapter.Commit(context.Background())
	assert.Nil(t, err)
}

func TestCommit_Error(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock).ShareClient(true))
	ses, _ := mt.Client.StartSession()
	adapter := &mongoAdapter{
		client: mt.Client,
		sess:   ses,
	}
	mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{})) // Simulate error
	err := adapter.Commit(context.Background())
	assert.Error(t, err)
}

func TestRollback_Success(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock).ShareClient(true))
	mt.AddMockResponses(mtest.CreateSuccessResponse())
	ses, _ := mt.Client.StartSession()
	adapter := &mongoAdapter{
		client: mt.Client,
		sess:   ses,
	}

	adapter.Begin(context.Background()) // Start transaction

	err := adapter.Rollback(context.Background())
	assert.Nil(t, err)
}

func TestRollback_NoSession(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock).ShareClient(true))

	adapter := &mongoAdapter{client: mt.Client} // No session set
	err := adapter.Rollback(context.Background())
	assert.Nil(t, err)
}

func TestRollback_Error(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock).ShareClient(true))
	ses, _ := mt.Client.StartSession()
	adapter := &mongoAdapter{
		client: mt.Client,
		sess:   ses,
	}
	mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{})) // Simulate error
	err := adapter.Rollback(context.Background())
	assert.Error(t, err)
}

func TestEnd(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock).ShareClient(true))

	// Mock session end
	mt.AddMockResponses(mtest.CreateSuccessResponse())
	ses, _ := mt.Client.StartSession()
	adapter := &mongoAdapter{
		client: mt.Client,
		sess:   ses,
	}

	adapter.End(context.Background())
	// No assertions needed, as EndSession doesn't return an error in the mock
}

func TestEnd_NoSession(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock).ShareClient(true))

	adapter := &mongoAdapter{client: mt.Client} // No session set

	// Calling End with no session should not panic
	assert.NotPanics(t, func() { adapter.End(context.Background()) })
}

func TestGetCurrent(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock).ShareClient(true))

	// Mock session end
	mt.AddMockResponses(mtest.CreateSuccessResponse())
	adapter := &mongoAdapter{
		client: mt.Client,
	}

	adapter.Begin(context.Background())

	assert.Equal(t, adapter.sessCtx, adapter.GetCurrent(context.Background()))

	adapter.End(context.Background())

	ctx := context.Background()
	assert.Equal(t, ctx, adapter.GetCurrent(ctx))
}
