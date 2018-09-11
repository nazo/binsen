package db

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

// NewDBConfig todo
type NewDBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

// NewDB todo
func NewDB(config *NewDBConfig) (*sql.DB, error) {
	var err error
	db, err = sql.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			config.Host,
			config.Port,
			config.User,
			config.Database,
			config.Password,
			config.SSLMode,
		),
	)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// InjectDB todo
func InjectDB() *sql.DB {
	return db
}
