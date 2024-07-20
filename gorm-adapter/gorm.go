package gormadapter

import (
	"context"

	"github.com/9ssi7/txn"
	"gorm.io/gorm"
)

type Gorm txn.Txn[*gorm.DB]

type gormTxn struct {
	tx  *gorm.DB
	cbs []txn.Callback[*gorm.DB]
}

func NewGorm(db *gorm.DB) txn.Txn[*gorm.DB] {
	return &gormTxn{
		tx: db,
	}
}

func (t *gormTxn) Add(cb txn.Callback[*gorm.DB]) {
	t.cbs = append(t.cbs, cb)
}

func (t *gormTxn) Transaction(_ context.Context) error {
	return t.tx.Transaction(func(tx *gorm.DB) error {
		for _, cb := range t.cbs {
			if err := cb(tx); err != nil {
				return err
			}
		}
		return nil
	})
}
