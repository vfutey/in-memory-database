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
	res, ok := db.Get(key1)
	require.False(t, ok)
	assert.Empty(t, res)

	db.Set(key1, value1)
	res, ok = db.Get(key1)
	require.True(t, ok)
	assert.Equal(t, value1, res)

	db.Delete(key1)
	res, ok = db.Get(key1)
	require.False(t, ok)
	assert.Empty(t, res)
}
