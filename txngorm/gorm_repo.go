package txngorm

import (
	"context"

	"github.com/9ssi7/txn"
	"gorm.io/gorm"
)

type txnRepo struct {
	tx *gorm.DB
	db *gorm.DB
}

// NewTxnRepo creates a new transaction-enabled repository for the given GORM database connection (*gorm.DB).
// This repository implements the Repo interface from the 9ssi7/txn package and provides all the necessary methods for transaction management with GORM.
func NewTxnRepo(db *gorm.DB) txn.Repo[*gorm.DB] {
	return &txnRepo{
		db: db,
	}
}

func (r *txnRepo) WithTxn(db *gorm.DB) {
	r.tx = db
}

func (r *txnRepo) GetCurrent(ctx context.Context) *gorm.DB {
	if r.tx != nil {
		return r.tx
	}
	return r.db.WithContext(ctx)
}

func (r *txnRepo) ClearTxn() {
	r.tx = nil
}
