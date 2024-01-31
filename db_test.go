package in_memory_database

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	key1   = "key1"
	value1 = "value1"
)

func TestSimple(t *testing.T) {
	db := InMemoryDatabase()
	assertAbsent(t, db, key1)

	db.Set(key1, value1)
	assertEqual(t, db, key1, value1)

	db.Delete(key1)
	assertAbsent(t, db, key1)
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
