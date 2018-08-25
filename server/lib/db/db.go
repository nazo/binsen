package db

import (
	"database/sql"
)

// NewDB todo
func NewDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=postgres port=5432 user=binsen dbname=binsen_dev password=binsen sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, nil
}
