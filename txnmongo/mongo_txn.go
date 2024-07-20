package txnmongo

import (
	"context"

	"github.com/9ssi7/txn"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Txr = context.Context

type mongoTxn struct {
	client *mongo.Client
	cbs    []txn.Callback[Txr]
	opts   []*options.SessionOptions
}

// New creates a new Mongo transaction object for the given MongoDB client.
// This object implements the Txn interface from 9ssi7/txn and provides the
// necessary methods to manage transactions within MongoDB.
func New(client *mongo.Client, opts ...*options.SessionOptions) txn.Txn[Txr] {
	return &mongoTxn{
		client: client,
		opts:   opts,
	}
}

func (t *mongoTxn) Add(cb txn.Callback[Txr]) {
	t.cbs = append(t.cbs, cb)
}

func (t *mongoTxn) Transaction(ctx context.Context) error {
	session, err := t.client.StartSession(t.opts...)
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)
	_, err = session.WithTransaction(ctx, func(sesctx mongo.SessionContext) (interface{}, error) {
		for _, cb := range t.cbs {
			if err := cb(sesctx); err != nil {
				return nil, err
			}
		}
		if err := session.CommitTransaction(ctx); err != nil {
			return nil, err
		}
		return nil, nil
	})
	return err
}
