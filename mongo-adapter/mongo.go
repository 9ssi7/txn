package mongoadapter

import (
	"context"

	"github.com/9ssi7/txn"
	"go.mongodb.org/mongo-driver/mongo"
)

type Mongo txn.Txn[mongo.SessionContext]

type mongoTxn struct {
	client *mongo.Client
	cbs    []txn.Callback[mongo.SessionContext]
}

func NewMongo(client *mongo.Client) Mongo {
	return &mongoTxn{
		client: client,
	}
}

func (t *mongoTxn) Add(cb txn.Callback[mongo.SessionContext]) {
	t.cbs = append(t.cbs, cb)
}

func (t *mongoTxn) Transaction(ctx context.Context) error {
	session, err := t.client.StartSession()
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
