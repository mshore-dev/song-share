package database

import (
	"database/sql"
	"log"

	_ "embed"

	_ "github.com/mattn/go-sqlite3"

	"song-share/database/sqlc"
)

func New(uri string) {

	db, err := sql.Open("sqlite3", uri)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	sqlc.New(db)
}
