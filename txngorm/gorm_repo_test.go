package txngorm

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getMockDB() *gorm.DB {
	db, _, _ := sqlmock.New()
	defer db.Close()

	gormDB, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	return gormDB
}

func TestTxnRepo_NewTxnRepo(t *testing.T) {
	db := getMockDB()
	repo := NewTxnRepo(db)
	assert.NotNil(t, repo)
	assert.NotEqual(t, db, repo.GetCurrent(context.Background()))
}

func TestTxnRepo_WithTxn(t *testing.T) {
	db := getMockDB()
	repo := NewTxnRepo(db)
	repo.WithTxn(db)
	assert.Equal(t, db, repo.GetCurrent(context.Background()))
}

func TestTxnRepo_GetCurrentDB(t *testing.T) {
	db := getMockDB()
	repo := NewTxnRepo(db)
	repo.WithTxn(db)
	assert.Equal(t, db, repo.GetCurrent(context.Background()))

	// Test without transaction
	repo.ClearTxn()
	assert.NotEqual(t, db, repo.GetCurrent(context.Background()))
}

func TestTxnRepo_ClearTxn(t *testing.T) {
	db := getMockDB()
	repo := txnRepo{
		db: db,
	}
	repo.WithTxn(db)
	repo.ClearTxn()
	assert.Nil(t, repo.tx)
	assert.NotEqual(t, db, repo.GetCurrent(context.Background()))
}
