package config

import (
	"os"

	"github.com/rodrigoenzohernandez/go-albums-service/internal/utils/logger"

	"github.com/joho/godotenv"
)

var log = logger.GetLogger("config")

// Load the environment variables from the .env file.
func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Error("Error loading .env file")
	}

}

// GetEnv returns the value of an environment variable or a fallback value if it doesn't exist.
func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
