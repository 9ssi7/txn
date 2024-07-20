package txnmongo

import (
	"context"

	"github.com/9ssi7/txn/tx"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MAdapter is the interface for interacting with MongoDB within a transaction.
// It extends the txn.Adapter interface to provide additional MongoDB-specific functionality.
type MAdapter interface {
	tx.Adapter

	// GetCurrent returns the current context.Context to use for executing MongoDB commands.
	// Depending on the transaction state, this may be the original context or a
	// mongo.SessionContext.
	GetCurrent(ctx context.Context) context.Context
}

// New creates a new MAdapter instance using the provided mongo.Client.
func New(client *mongo.Client) MAdapter {
	return &mongoAdapter{client: client}
}

type mongoAdapter struct {
	client    *mongo.Client
	sess      mongo.Session
	sessCtx   mongo.SessionContext
	sesOption *options.SessionOptions
	txOption  *options.TransactionOptions
}

func (a *mongoAdapter) Begin(ctx context.Context) error {
	ses, err := a.client.StartSession(a.sesOption)
	if err != nil {
		return err
	}
	a.sess = ses
	a.sessCtx = mongo.NewSessionContext(ctx, a.sess)
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
		a.sess = nil
		a.sessCtx = nil
	}
}

func (a *mongoAdapter) GetCurrent(ctx context.Context) context.Context {
	if a.sessCtx != nil {
		return a.sessCtx
	}
	return ctx
}
