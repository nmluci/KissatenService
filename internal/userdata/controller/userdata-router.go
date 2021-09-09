package controller

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func RegisterUserdataSubrouter(r *mux.Router, db *sql.DB) {
	r.HandleFunc("/api/user/", GetAllUserdataController()).Methods("GET")
	r.HandleFunc("/api/user/{username}", GetUserDataController()).Methods("GET")
	r.HandleFunc("/api/user/{username}/register", RegisterNewUser()).Methods("POST")
}