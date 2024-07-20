package txnmongo

import (
	"context"

	"github.com/9ssi7/txn"
	"go.mongodb.org/mongo-driver/mongo"
)

// RepoTxr is a transaction context for MongoDB.
// This type alias is used to define the transaction context for MongoDB.
type RepoTxr context.Context

// Repo is a transaction-enabled repository for MongoDB.
// This type alias is used to define the repository methods that work with
type Repo txn.Repo[RepoTxr]

type txnRepo struct {
	client     *mongo.Client
	sessionCtx RepoTxr
}

// NewTxnRepo creates a new transaction-enabled repository for the given
// MongoDB client. This repository implements the Repo interface from 9ssi7/txn
// and provides the necessary methods to work with transactions in MongoDB.
func NewTxnRepo(client *mongo.Client) txn.Repo[RepoTxr] {
	return &txnRepo{
		client: client,
	}
}

func (r *txnRepo) WithTxn(sessionCtx RepoTxr) {
	r.sessionCtx = sessionCtx
}

func (r *txnRepo) GetCurrent(ctx context.Context) RepoTxr {
	if r.sessionCtx == nil {
		return ctx
	}
	return r.sessionCtx
}

func (r *txnRepo) ClearTxn() {
	r.sessionCtx = nil
}
