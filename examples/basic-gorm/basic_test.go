package main

import (
	"context"
	"testing"

	"github.com/9ssi7/txn/txngorm"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestBasic(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	gormDB, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	txn := txngorm.New(gormDB)

	repo := NewCustomRepo(gormDB)

	mock.ExpectBegin()
	mock.ExpectExec("SELECT 1").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	runGenericService(context.Background(), txn, repo)
}
