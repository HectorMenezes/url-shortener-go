package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	result := Hash(7, "https://github.com")
	assert.Equal(t, "2931944", result, "Should be equal hash.")
}

func TestGetEnv(t *testing.T) {
	os.Setenv("FOO", "1")
	result_foo := GetEnv("FOO", "2")
	result_bar := GetEnv("BAR", "3")
	assert.Equal(t, "1", result_foo, "Should be equal to 1.")
	assert.Equal(t, "3", result_bar, "Should be equal to 3.")
}
