package txngorm

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Mock Callback for testing
func mockCallback(db *gorm.DB) error {
	return db.Exec("SELECT 1").Error // Simulate a simple DB interaction
}

func TestNewGorm(t *testing.T) {
	db, _, _ := sqlmock.New()
	defer db.Close()

	gormDB, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	txn := NewGorm(gormDB)

	if txn == nil {
		t.Fatal("NewGorm returned nil")
	}
}

func TestGormTxn_AddWithState(t *testing.T) {
	type state struct {
		Id   int
		Name string
	}
	db, mock, _ := sqlmock.New()
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("SELECT 1").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	gormDB, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	txn := NewGorm(gormDB)
	var s state
	txn.Add(func(db *gorm.DB) error {
		s = state{Id: 1, Name: "test"}
		return nil
	})

	txn.Add(func(db *gorm.DB) error {
		if s.Id != 1 || s.Name != "test" {
			t.Fatal("State not passed between callbacks")
		}
		return nil
	})

	txn.Add(mockCallback)

	err := txn.Transaction(context.Background())
	if err != nil {
		t.Fatalf("Transaction failed: %v", err)
	}
}

func TestGormTxn_Add(t *testing.T) {
	db, _, _ := sqlmock.New()
	defer db.Close()

	gormDB, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	txn := NewGorm(gormDB)

	txn.Add(mockCallback)
	if len(txn.(*gormTxn).cbs) != 1 {
		t.Fatal("Add did not append callback")
	}
}

func TestGormTxn_Transaction_Success(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("SELECT 1").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	gormDB, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	txn := NewGorm(gormDB)
	txn.Add(mockCallback)

	err := txn.Transaction(context.Background())
	if err != nil {
		t.Fatalf("Transaction failed: %v", err)
	}
}

func TestGormTxn_Transaction_CallbackError(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("SELECT 1").WillReturnError(errors.New("DB error"))
	mock.ExpectRollback()

	gormDB, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	txn := NewGorm(gormDB)
	txn.Add(mockCallback)

	err := txn.Transaction(context.Background())
	if err == nil {
		t.Fatal("Expected Transaction to fail, but it succeeded")
	}
}

func TestGormTxn_Transaction_BeginError(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	mock.ExpectBegin().WillReturnError(errors.New("Begin error"))

	gormDB, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	txn := NewGorm(gormDB)

	err := txn.Transaction(context.Background())
	if err == nil {
		t.Fatal("Expected Transaction to fail, but it succeeded")
	}
}
