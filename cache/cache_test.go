package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCache(t *testing.T) {
	assert.Equal(t, (*allCache)(nil), GetCache(), "Should be nil at first")
	Start()
	assert.NotEqual(t, (*allCache)(nil), GetCache(), "Shouldn't be nil anymore")
}

