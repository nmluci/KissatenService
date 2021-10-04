package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	dbm "github.com/nmluci/KissatenService/database/models"
	"github.com/nmluci/KissatenService/database/router"
)

func InitializeDatabase() *sql.DB {
	if db, err := sql.Open("sqlite3", "./storage/database.db"); err != nil {
		log.Fatalf("Database Error: %s", err.Error())
		return nil
	} else {
		return db
	}
}

func main() {
	r := mux.NewRouter()

	db := &dbm.DatabaseModel{}
	db.DB = InitializeDatabase()
	router.Router(r, db)

	log.Fatal(http.ListenAndServe("localhost:8081", r))
}
