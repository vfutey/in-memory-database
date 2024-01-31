package in_memory_database

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	key1   = "key1"
	value1 = "value1"
	value2 = "value2"
)

func TestSimple(t *testing.T) {
	db := InMemoryDatabase()
	assertAbsent(t, db, key1)

	db.Set(key1, value1)
	assertEqual(t, db, key1, value1)

	db.Delete(key1)
	assertAbsent(t, db, key1)
}

func TestTransactionCommit(t *testing.T) {
	// Example 1 for commit a transaction
	db := InMemoryDatabase()
	db.Set(key1, value1)
	db.StartTransaction()
	db.Set(key1, value2)
	db.Commit()
	assertEqual(t, db, key1, value2)
}

func TestTransactionRollback(t *testing.T) {
	// Example 2 for roll_back().
	db := InMemoryDatabase()
	db.Set(key1, value1)
	db.StartTransaction()
	assertEqual(t, db, key1, value1)
	db.Set(key1, value2)
	assertEqual(t, db, key1, value2)
	db.Rollback()
	assertEqual(t, db, key1, value1)
}

func TestNestedTransactionCommit(t *testing.T) {
	// Example 3 for nested transactions
	db := InMemoryDatabase()
	db.Set(key1, value1)
	db.StartTransaction()
	db.Set(key1, value2)
	assertEqual(t, db, key1, value2)
	db.StartTransaction()
	assertEqual(t, db, key1, value2)
	db.Delete(key1)
	db.Commit()
	assertAbsent(t, db, key1)
	db.Commit()
	assertAbsent(t, db, key1)
}

func TestNestedTransactionRollback(t *testing.T) {
	// Example 4 for nested transactions with rollback()
	db := InMemoryDatabase()
	db.Set(key1, value1)
	db.StartTransaction()
	db.Set(key1, value2)
	assertEqual(t, db, key1, value2)
	db.StartTransaction()
	assertEqual(t, db, key1, value2)
	db.Delete(key1)
	db.Rollback()
	assertEqual(t, db, key1, value2)
	db.Commit()
	assertEqual(t, db, key1, value2)
}

func assertEqual(t *testing.T, db DB, key, value string) {
	t.Helper()
	res, ok := db.Get(key)
	require.True(t, ok)
	assert.Equal(t, value, res)
}

func assertAbsent(t *testing.T, db DB, key string) {
	t.Helper()
	res, ok := db.Get(key)
	require.False(t, ok)
	assert.Empty(t, res)
}
