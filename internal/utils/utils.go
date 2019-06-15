/*
Package utils exports miscellaneous utility functions
*/
package utils

import (
	"fmt"
	"os"
	"strconv"
)

// GetEnv returns environment variable if present otherwise returns fallback
func GetEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// GetEnvAsInt returns environment variable as integer if present otherwise returns fallback.
// Error is returned if value can not be converted to integer.
func GetEnvAsInt(key string, fallback int) (int, error) {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.Atoi(value)
    if err != nil {
        return 0, fmt.Errorf("Unabled to convert environment variable '%v' to integer", key)
		}
		return i, nil
	}
	return fallback, nil
}

// GetEnvOrPanic returns environment variable if present otherwise panics
func GetEnvOrPanic(key string) string {
	value, ok := os.LookupEnv(key);
	if !ok {
		panic(fmt.Sprintf("Missing required environment variable '%v'\n", key))
	}
	return value
}
