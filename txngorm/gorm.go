package txngorm

import (
	"context"

	"github.com/9ssi7/txn"
	"gorm.io/gorm"
)

type gormAdapter struct {
	db *gorm.DB
	tx *gorm.DB
}

func New(db *gorm.DB) txn.Adapter {
	return &gormAdapter{db: db}
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
