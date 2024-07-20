package main

import (
	"context"
	"fmt"

	"github.com/9ssi7/txn"
	"github.com/9ssi7/txnmongo"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repo[Txr any] interface {
	txn.Repo[Txr]
	Insert(ctx context.Context, id string) error
}

type custom struct {
	ID string `bson:"_id"`
}

type customRepo struct {
	txn.Repo[txnmongo.Txr]
	collection *mongo.Collection
}

func NewCustomRepo(client *mongo.Client) Repo[txnmongo.Txr] {
	return &customRepo{
		Repo:       txnmongo.NewTxnRepo(client),
		collection: client.Database("test").Collection("test"),
	}
}

func (r *customRepo) Insert(ctx context.Context, id string) error {
	return r.collection.FindOne(r.Repo.GetCurrent(ctx), custom{ID: id}).Err()
}

func runGenericService[Txr any](ctx context.Context, txn txn.Txn[Txr], repo Repo[Txr]) {
	txn.Add(func(d Txr) error {
		repo.WithTxn(d)
		return repo.Insert(ctx, "1")
	})
	if err := txn.Transaction(ctx); err != nil {
		repo.ClearTxn()
		fmt.Println(err)
	}
}
