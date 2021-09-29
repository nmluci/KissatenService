package database

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitializeDatabase(ctx context.Context) *sql.DB {
	if db, err := sql.Open("sqlite3", "./pkg/database/database.db"); err != nil {
		log.Fatalf("Database Error: %s", err.Error())
		return nil
	} else {
		return db
	}
}
