package repository

import (
	"database/sql"
	"fmt"

	"github.com/rodrigoenzohernandez/go-albums-service/config"
	"github.com/rodrigoenzohernandez/go-albums-service/internal/utils/logger"

	_ "github.com/lib/pq"
)

var log = logger.GetLogger("repository")

func Connect() *sql.DB {

	host := config.GetEnv("DB_HOST", "localhost")
	user := config.GetEnv("DB_USER", "postgres")
	dbName := config.GetEnv("DB_NAME", "dev")
	password := config.GetEnv("DB_PASSWORD", "p4ssw0rd.db")
	connectTimeout := config.GetEnv("DB_CONNECT_TIMEOUT", "5")
	SSLMode := config.GetEnv("DB_SSL_MODE", "enable")

	connStr := fmt.Sprintf("host=%s user=%s dbname=%s password=%s connect_timeout=%s sslmode=%s",
		host, user, dbName, password, connectTimeout, SSLMode)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Error("Error connecting to the database.")
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		log.Error("Error reaching to the database.")
		panic(err)

	}

	log.Info("Successfully connected to the database.")

	return db
}

func Disconnect(db *sql.DB) {
	if err := db.Close(); err != nil {
		log.Error("Error disconnecting from the database.")
		panic(err)

	}
}
