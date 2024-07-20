package txnmongo

import (
	"context"

	"github.com/9ssi7/txn"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoAdapter struct {
	client    *mongo.Client
	sess      mongo.Session
	sesOption *options.SessionOptions
	txOption  *options.TransactionOptions
}

func New(client *mongo.Client) txn.Adapter {
	return &mongoAdapter{client: client}
}

func (a *mongoAdapter) Begin(ctx context.Context) error {
	ses, err := a.client.StartSession(a.sesOption)
	if err != nil {
		return err
	}
	a.sess = ses
	return a.sess.StartTransaction(a.txOption)
}

func (a *mongoAdapter) Commit(ctx context.Context) error {
	if a.sess == nil {
		return nil
	}
	return a.sess.CommitTransaction(ctx)
}

func (a *mongoAdapter) Rollback(ctx context.Context) error {
	if a.sess == nil {
		return nil
	}
	return a.sess.AbortTransaction(ctx)
}

func (a *mongoAdapter) End(ctx context.Context) {
	if a.sess != nil {
		a.sess.EndSession(ctx)
	}
}
