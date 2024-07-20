package txnmongo

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestNewTxnRepo(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	repo := NewTxnRepo(mt.Client)
	if repo == nil {
		t.Fatal("NewTxnRepo returned nil")
	}
}

func TestTxnRepo_WithTxn(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock).ShareClient(true))

	repo := NewTxnRepo(mt.Client)

	session, err := mt.Client.StartSession()
	if err != nil {
		t.Fatalf("Failed to start session: %v", err)
	}
	mockSessionCtx := mongo.NewSessionContext(context.Background(), session)

	repo.WithTxn(mockSessionCtx)

	if repo.(*txnRepo).sessionCtx != mockSessionCtx {
		t.Fatal("WithTxn did not set sessionCtx correctly")
	}
}

func TestTxnRepo_GetCurrent_WithSession(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock).ShareClient(true))

	repo := NewTxnRepo(mt.Client)

	session, err := mt.Client.StartSession()
	if err != nil {
		t.Fatalf("Failed to start session: %v", err)
	}
	mockSessionCtx := mongo.NewSessionContext(context.Background(), session)
	repo.WithTxn(mockSessionCtx)

	currentCtx := repo.GetCurrent(context.Background())

	if currentCtx != mockSessionCtx {
		t.Fatal("GetCurrent did not return sessionCtx when set")
	}
}

func TestTxnRepo_GetCurrent_WithoutSession(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock).ShareClient(true))

	repo := NewTxnRepo(mt.Client)

	ctx := context.Background()
	currentCtx := repo.GetCurrent(ctx)

	if currentCtx != ctx {
		t.Fatal("GetCurrent did not return original context when sessionCtx is nil")
	}
}

func TestTxnRepo_ClearTxn(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock).ShareClient(true))

	repo := NewTxnRepo(mt.Client)

	session, err := mt.Client.StartSession()
	if err != nil {
		t.Fatalf("Failed to start session: %v", err)
	}
	mockSessionCtx := mongo.NewSessionContext(context.Background(), session)
	repo.WithTxn(mockSessionCtx)

	repo.ClearTxn()

	if repo.(*txnRepo).sessionCtx != nil {
		t.Fatal("ClearTxn did not clear sessionCtx")
	}
}
