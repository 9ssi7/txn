package main

import (
	"context"
	"testing"

	"github.com/9ssi7/txnmongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestBasic(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock).ShareClient(true))

	mt.AddMockResponses(mtest.CreateSuccessResponse(), mtest.CreateSuccessResponse())
	client := mt.Client
	txn := txnmongo.New(client)

	repo := NewCustomRepo(client)

	runGenericService(context.Background(), txn, repo)
}
