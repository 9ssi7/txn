package txngorm

import (
	"context"

	"github.com/9ssi7/txn"
	"gorm.io/gorm"
)

// Gorm is a type alias for the txn gorm implementation.
type Gorm txn.Txn[*gorm.DB]

type gormTxn struct {
	tx  *gorm.DB
	cbs []txn.Callback[*gorm.DB]
}

// NewGorm creates a new Gorm transaction object for the given GORM database connection (*gorm.DB).
// This object implements the Txn interface from the 9ssi7/txn package and provides all the necessary methods for transaction management with GORM.
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
				tx.Rollback()
				return err
			}
		}
		return nil
	})
}
