package controller

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/nmluci/KissatenService/internal/userdata/models"
)

func RegisterUserdataSubrouter(r *mux.Router, db *sql.DB) {
	um := &models.UserModel{DB: db}
	r.HandleFunc("/api/user", GetAllUserdataController(um)).Methods("GET")
	r.HandleFunc("/api/user/{username}", GetUserDataController(um)).Methods("GET")
	r.HandleFunc("/api/user/register", RegisterNewUser(um)).Methods("POST")
}
