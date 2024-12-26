package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func MustNew(dsn string) *sql.DB {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		panic(err)
	}
	return db
}
