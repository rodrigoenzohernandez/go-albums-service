package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLogger(t *testing.T) {
	logger := GetLogger("database")
	assert.NotNil(t, logger, "GetLogger returned nil")
}

func TestGetLoggerSameInstance(t *testing.T) {
	logger1 := GetLogger("database")
	logger2 := GetLogger("database")
	assert.Equal(t, logger1, logger2, "Expected GetLogger to return the same instance for 'database'")
}

func TestGetLoggerDifferentInstance(t *testing.T) {
	logger1 := GetLogger("database")
	logger2 := GetLogger("server")
	assert.NotEqual(t, logger1, logger2, "Expected GetLogger to return different instances for 'database' and 'server'")
}
