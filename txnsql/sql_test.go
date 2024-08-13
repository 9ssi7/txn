package txnsql

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	adapter := New(db)
	assert.NotNil(t, adapter)
	assert.IsType(t, &sqlAdapter{}, adapter)
}

func TestSqlAdapter_Begin(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	t.Run("Begin success", func(t *testing.T) {
		adapter := &sqlAdapter{db: db}
		mock.ExpectBegin()

		err := adapter.Begin(context.Background())
		assert.Nil(t, err)
		assert.NotNil(t, adapter.tx)
	})

	t.Run("Begin error", func(t *testing.T) {
		adapter := &sqlAdapter{db: db}
		mock.ExpectBegin().WillReturnError(errors.New("begin failed"))

		err := adapter.Begin(context.Background())
		assert.NotNil(t, err)
		assert.Nil(t, adapter.tx)
	})
}

func TestSqlAdapter_Commit(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	adapter := &sqlAdapter{db: db}

	t.Run("Commit success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectCommit()

		adapter.Begin(context.Background())
		err := adapter.Commit(context.Background())
		assert.Nil(t, err)
		assert.Nil(t, adapter.tx)
	})

	t.Run("Commit error", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectCommit().WillReturnError(errors.New("commit failed"))

		adapter.Begin(context.Background())
		err := adapter.Commit(context.Background())
		assert.NotNil(t, err)
		assert.Nil(t, adapter.tx)
	})

	t.Run("Commit no tx", func(t *testing.T) {
		err := adapter.Commit(context.Background())
		assert.Nil(t, err)
		assert.Nil(t, adapter.tx)
	})
}

func TestSqlAdapter_Rollback(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	adapter := &sqlAdapter{db: db}

	t.Run("Rollback success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectRollback()

		adapter.Begin(context.Background())
		err := adapter.Rollback(context.Background())
		assert.Nil(t, err)
		assert.Nil(t, adapter.tx)
	})

	t.Run("Rollback error", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectRollback().WillReturnError(errors.New("rollback failed"))

		adapter.Begin(context.Background())
		err := adapter.Rollback(context.Background())
		assert.NotNil(t, err)
		assert.Nil(t, adapter.tx)
	})

	t.Run("Rollback no tx", func(t *testing.T) {
		err := adapter.Rollback(context.Background())
		assert.Nil(t, err)
		assert.Nil(t, adapter.tx)
	})
}

func TestSqlAdapter_End(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	adapter := &sqlAdapter{db: db}

	t.Run("End", func(t *testing.T) {
		mock.ExpectBegin()

		adapter.Begin(context.Background())
		adapter.End(context.Background())
		assert.Nil(t, adapter.tx)
	})

	t.Run("End no tx", func(t *testing.T) {
		adapter.End(context.Background())
		assert.Nil(t, adapter.tx)
	})
}

func TestSqlAdapter_Tx(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	t.Run("Tx is not nil", func(t *testing.T) {
		adapter := &sqlAdapter{db: db}
		mock.ExpectBegin()

		adapter.Begin(context.Background())
		assert.True(t, adapter.Tx() != nil)
	})

	t.Run("Tx is nil", func(t *testing.T) {
		adapter := &sqlAdapter{db: db}
		assert.True(t, adapter.Tx() == nil)
	})
}
