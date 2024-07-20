package txnmongo

import (
	"context"

	"github.com/9ssi7/txn"
	"go.mongodb.org/mongo-driver/mongo"
)

type txnRepo struct {
	client     *mongo.Client
	sessionCtx Txr
}

// NewTxnRepo creates a new transaction-enabled repository for the given
// MongoDB client. This repository implements the Repo interface from 9ssi7/txn
// and provides the necessary methods to work with transactions in MongoDB.
func NewTxnRepo(client *mongo.Client) txn.Repo[Txr] {
	return &txnRepo{
		client: client,
	}
}

func (r *txnRepo) WithTxn(sessionCtx Txr) {
	r.sessionCtx = sessionCtx
}

func (r *txnRepo) GetCurrent(ctx context.Context) Txr {
	if r.sessionCtx == nil {
		return ctx
	}
	return r.sessionCtx
}

func (r *txnRepo) ClearTxn() {
	r.sessionCtx = nil
}
