package utils

import (
	"hash/fnv"
	"os"
	"strconv"
)

// Hash returns hash with specified size from string.
func Hash(length int, s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))
	return strconv.FormatUint(uint64(h.Sum32()), 10)[:length]
}

// GetEnv returns the environment variable.
// If it isn't found, return fallback.
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
