package tx

import (
	"context"
	"errors"
	"testing"
)

// MockAdapter implements the Adapter interface for testing
type MockAdapter struct {
	beginError    error
	commitError   error
	rollbackError error
}

func (ma *MockAdapter) Begin(ctx context.Context) error    { return ma.beginError }
func (ma *MockAdapter) Commit(ctx context.Context) error   { return ma.commitError }
func (ma *MockAdapter) Rollback(ctx context.Context) error { return ma.rollbackError }
func (ma *MockAdapter) End(ctx context.Context)            {}

func TestNew(t *testing.T) {
	tx := New()
	if tx == nil {
		t.Fatal("New returned nil")
	}
}

func TestRegister(t *testing.T) {
	tx := New()
	adapter := &MockAdapter{}

	tx.Register(adapter)

	if len(tx.(*txn).adapters) != 1 || tx.(*txn).adapters[0] != adapter {
		t.Fatal("Register failed to add adapter")
	}
}

func TestBegin_Success(t *testing.T) {
	tx := New()
	tx.Register(&MockAdapter{})
	tx.Register(&MockAdapter{})

	err := tx.Begin(context.Background())

	if err != nil {
		t.Fatalf("Begin failed unexpectedly: %v", err)
	}
}

func TestBegin_Failure(t *testing.T) {
	tx := New()
	tx.Register(&MockAdapter{beginError: errors.New("begin error")})

	err := tx.Begin(context.Background())

	if err == nil {
		t.Fatal("Expected Begin to fail, but it succeeded")
	}
}

func TestCommit_Success(t *testing.T) {
	tx := New()
	tx.Register(&MockAdapter{})
	tx.Register(&MockAdapter{})

	err := tx.Commit(context.Background())

	if err != nil {
		t.Fatalf("Commit failed unexpectedly: %v", err)
	}
}

func TestCommit_Failure(t *testing.T) {
	tx := New()
	tx.Register(&MockAdapter{commitError: errors.New("commit error")})

	err := tx.Commit(context.Background())

	if err == nil {
		t.Fatal("Expected Commit to fail, but it succeeded")
	}
}

func TestRollback_Success(t *testing.T) {
	tx := New()
	tx.Register(&MockAdapter{})
	tx.Register(&MockAdapter{})

	err := tx.Rollback(context.Background())

	if err != nil {
		t.Fatalf("Rollback failed unexpectedly: %v", err)
	}
}

func TestRollback_Failure(t *testing.T) {
	tx := New()
	tx.Register(&MockAdapter{rollbackError: errors.New("rollback error")})

	err := tx.Rollback(context.Background())

	if err == nil {
		t.Fatal("Expected Rollback to fail, but it succeeded")
	}
}

func TestCancel(t *testing.T) {
	tx := New()
	adapter := &MockAdapter{} // Create a new instance to capture the rollback call
	tx.Register(adapter)

	tx.Cancel(context.Background())

	// Ideally, you'd assert that the adapter's Rollback method was called here.
	// Since it's a mock, you can add a flag to MockAdapter and check if it's set after Cancel.
}
