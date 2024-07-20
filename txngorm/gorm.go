package txngorm

import (
	"context"

	"github.com/9ssi7/txn"
	"gorm.io/gorm"
)

// GAdapter is the interface for interacting with GORM within a transaction.
// It extends the txn.Adapter interface to provide additional GORM-specific functionality.
type GAdapter interface {
	txn.Adapter

	// GetCurrent returns the current *gorm.DB instance to use for executing GORM operations.
	// Depending on the transaction state, this may be the original db instance or a transaction object.
	GetCurrent(ctx context.Context) *gorm.DB
}

// New creates a new GAdapter instance using the provided *gorm.DB.
func New(db *gorm.DB) GAdapter {
	return &gormAdapter{db: db}
}

type gormAdapter struct {
	db *gorm.DB
	tx *gorm.DB
}

func (a *gormAdapter) Begin(ctx context.Context) error {
	tx := a.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}
	a.tx = tx
	return nil
}

func (a *gormAdapter) Commit(_ context.Context) error {
	if a.tx == nil {
		return nil
	}
	err := a.tx.Commit().Error
	a.tx = nil
	return err
}

func (a *gormAdapter) Rollback(_ context.Context) error {
	if a.tx == nil {
		return nil
	}
	err := a.tx.Rollback().Error
	a.tx = nil
	return err
}

func (a *gormAdapter) End(_ context.Context) {
	if a.tx != nil {
		a.tx = nil
	}
}

func (a *gormAdapter) GetCurrent(ctx context.Context) *gorm.DB {
	if a.tx != nil {
		return a.tx
	}
	return a.db
}
