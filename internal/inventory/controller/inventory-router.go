package controller

import (
	"database/sql"

	"github.com/gorilla/mux"
	models "github.com/nmluci/KissatenService/internal/inventory/model"
)

func RegisterInventorySubrouter(r *mux.Router, db *sql.DB) {
	im := &models.InventoryModel{DB: db}
	r.HandleFunc("/api/inventory/", GetAllItemController(im)).Methods("GET")
	r.HandleFunc("/api/inventory/", InsertItemController(im)).Methods("POST")
	// Will be changed soon, into a queries-like params
	r.HandleFunc("/api/inventory/name/{itemName}", GetItemByNameController(im)).Methods("GET")
	r.HandleFunc("/api/inventory/id/{itemId}", GetItemController(im)).Methods("GET")
	r.HandleFunc("/api/inventory/id/{itemId}", UpdateItemController(im)).Methods("POST")
	r.HandleFunc("/api/inventory/id/{itemId}", RemoveItemController(im)).Methods("DELETE")
}
