package mongoadapter

import (
	"context"
	"errors"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestNewMongo(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	client := mt.Client
	txn := NewMongo(client)

	if txn == nil {
		t.Fatal("NewMongo returned nil")
	}
}

func TestMongoTxn_Add(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	//	defer mt.Close()

	client := mt.Client
	txn := NewMongo(client)

	txn.Add(func(sc mongo.SessionContext) error { return nil }) // Mock callback

	if len(txn.(*mongoTxn).cbs) != 1 {
		t.Fatal("Add did not append callback")
	}
}

func TestMongoTxn_Transaction_Success(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock).ShareClient(true))
	//defer mt.Close()

	mt.AddMockResponses(mtest.CreateSuccessResponse(), mtest.CreateSuccessResponse())
	client := mt.Client
	txn := NewMongo(client)
	txn.Add(func(sc mongo.SessionContext) error {
		_, err := sc.Client().Database("test").Collection("test").InsertOne(sc, bson.E{Key: "x", Value: 1})
		return err
	})

	err := txn.Transaction(context.Background())
	if err != nil {
		t.Fatalf("Transaction failed: %v", err)
	}
}

func TestMongoTxn_Transaction_CallbackError(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock).ShareClient(true))
	//	defer mt.Close()

	mt.AddMockResponses(mtest.CreateSuccessResponse()) // Simulate session start success
	client := mt.Client
	txn := NewMongo(client)

	txn.Add(func(sc mongo.SessionContext) error {
		return errors.New("Callback error")
	})

	err := txn.Transaction(context.Background())
	if err == nil {
		t.Fatal("Expected Transaction to fail, but it succeeded")
	}
}

func TestMongoTxn_Transaction_CommitError(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock).ShareClient(true))
	//defer mt.Close()

	mt.AddMockResponses(mtest.CreateSuccessResponse(), mtest.CreateCommandErrorResponse(mtest.CommandError{
		Code: 123,
	})) // Simulate commit failure
	client := mt.Client
	txn := NewMongo(client)
	txn.Add(func(sc mongo.SessionContext) error {
		_, err := sc.Client().Database("test").Collection("test").InsertOne(sc, bson.E{Key: "x", Value: 1})
		return err
	})

	err := txn.Transaction(context.Background())
	if err == nil {
		t.Fatal("Expected Transaction to fail, but it succeeded")
	}
}
