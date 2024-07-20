package txnmongo

import (
	"context"

	"github.com/9ssi7/txn"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mongo is a transaction context for MongoDB.
// This type alias is used to define the transaction context for MongoDB.
type Mongo txn.Txn[mongo.SessionContext]

type mongoTxn struct {
	client *mongo.Client
	cbs    []txn.Callback[mongo.SessionContext]
	opts   []*options.SessionOptions
}

// NewMongo creates a new Mongo transaction object for the given MongoDB client.
// This object implements the Txn interface from 9ssi7/txn and provides the
// necessary methods to manage transactions within MongoDB.
func NewMongo(client *mongo.Client, opts ...*options.SessionOptions) Mongo {
	return &mongoTxn{
		client: client,
		opts:   opts,
	}
}

func (t *mongoTxn) Add(cb txn.Callback[mongo.SessionContext]) {
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
