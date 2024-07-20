package txnmongo

import (
	"context"
	"errors"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestNew(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	client := mt.Client
	txn := New(client)

	if txn == nil {
		t.Fatal("New returned nil")
	}
}

func TestMongoTxn_Add(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	client := mt.Client
	txn := New(client)

	txn.Add(func(sc Txr) error { return nil })

	if len(txn.(*mongoTxn).cbs) != 1 {
		t.Fatal("Add did not append callback")
	}
}

func TestMongoTxn_Transaction_Success(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock).ShareClient(true))

	mt.AddMockResponses(mtest.CreateSuccessResponse(), mtest.CreateSuccessResponse())
	client := mt.Client
	txn := New(client)
	col := client.Database("test").Collection("test")
	txn.Add(func(sc Txr) error {
		_, err := col.InsertOne(sc, bson.E{Key: "x", Value: 1})
		return err
	})

	err := txn.Transaction(context.Background())
	if err != nil {
		t.Fatalf("Transaction failed: %v", err)
	}
}

func TestMongoTxn_Transaction_CallbackError(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock).ShareClient(true))

	mt.AddMockResponses(mtest.CreateSuccessResponse())
	client := mt.Client
	txn := New(client)

	txn.Add(func(sc Txr) error {
		return errors.New("Callback error")
	})

	err := txn.Transaction(context.Background())
	if err == nil {
		t.Fatal("Expected Transaction to fail, but it succeeded")
	}
}

func TestMongoTxn_Transaction_CommitError(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock).ShareClient(true))

	mt.AddMockResponses(mtest.CreateSuccessResponse(), mtest.CreateCommandErrorResponse(mtest.CommandError{
		Code: 123,
	}))
	client := mt.Client
	txn := New(client)
	col := client.Database("test").Collection("test")
	txn.Add(func(sc Txr) error {
		_, err := col.InsertOne(sc, bson.E{Key: "x", Value: 1})
		return err
	})

	err := txn.Transaction(context.Background())
	if err == nil {
		t.Fatal("Expected Transaction to fail, but it succeeded")
	}
}

func TestMongoTxn_Transaction_StartError(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("Expected Transaction to panic, but it did not")
		}
	}()

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock).ShareClient(true))

	mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
		Code: 123,
	}))
	client := mt.Client
	bl := true
	bla := true
	txn := New(client, &options.SessionOptions{
		Snapshot:          &bl,
		CausalConsistency: &bla,
	})
	col := client.Database("test").Collection("test")

	txn.Add(func(sc Txr) error {
		_, err := col.InsertOne(sc, bson.E{Key: "x", Value: 1})
		return err
	})

	err := txn.Transaction(context.Background())
	if err == nil {
		t.Fatal("Expected Transaction to fail, but it succeeded")
	}

}
