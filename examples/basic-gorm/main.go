package main

import (
	"context"
	"fmt"

	"github.com/9ssi7/txn"
	"github.com/9ssi7/txn/txngorm"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repo[Txr txn.ITxr] interface {
	txn.Repo[Txr]
	Select(ctx context.Context) error
}

type customRepo struct {
	txn.Repo[*gorm.DB]
}

func NewCustomRepo(db *gorm.DB) Repo[*gorm.DB] {
	return &customRepo{
		Repo: txngorm.NewTxnRepo(db),
	}
}

func (r *customRepo) Select(ctx context.Context) error {
	return r.GetCurrent(ctx).Exec("SELECT 1").Error
}

func main() {
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

func runGenericService[Txr txn.ITxr](ctx context.Context, txn txn.Txn[Txr], repo Repo[Txr]) {
	txn.Add(func(d Txr) error {
		repo.WithTxn(d)
		return repo.Select(ctx)
	})
	if err := txn.Transaction(ctx); err != nil {
		repo.ClearTxn()
		fmt.Println(err)
	}
}
