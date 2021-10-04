package router

import (
	"github.com/gorilla/mux"
	"github.com/nmluci/KissatenService/database/controller"
	"github.com/nmluci/KissatenService/database/models"
)

func Router(r *mux.Router, dbm *models.DatabaseModel) {

	r.HandleFunc("/api/database/{service}", controller.GetItem(*dbm)).Methods("GET")
	r.HandleFunc("/api/database/{service}", controller.PostItem(*dbm)).Methods("POST")
	r.HandleFunc("/api/database", controller.RegisterService(dbm)).Methods("POST")
}
