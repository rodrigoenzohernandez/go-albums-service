package config_test

import (
	"os"
	"testing"

	"github.com/rodrigoenzohernandez/web-service-gin/config"
	"github.com/stretchr/testify/assert"
)

func TestGetEnv_ExistingKey(t *testing.T) {
	const testKey = "DB_USER"
	const testValue = "postgresAdmin"

	os.Setenv(testKey, testValue)
	defer os.Unsetenv(testKey)

	val := config.GetEnv(testKey, "")
	assert.Equal(t, testValue, val, "they should be equal")
}

func TestGetEnv_NonExistingKey(t *testing.T) {
	const fallbackValue = "postgres"

	val := config.GetEnv("NON_EXISTENT_KEY", fallbackValue)
	assert.Equal(t, fallbackValue, val, "fallback value should be returned when the key does not exist")
}
