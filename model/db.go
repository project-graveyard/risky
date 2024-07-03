package model

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// InitDB initialises a sqlite database
func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./risky.db")
	if err != nil {
		return nil, err
	}

	return db, nil
}
