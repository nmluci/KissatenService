package controller

import "github.com/gorilla/mux"

func RegisterInventorySubrouter(r *mux.Router) {
	r.HandleFunc("/api/inventory/", GetAllItemController()).Methods("GET")
	r.HandleFunc("/api/inventory/", InsertItemController()).Methods("POST")
	r.HandleFunc("/api/inventory/{itemId}", GetItemController()).Methods("GET")
	r.HandleFunc("/api/inventory/{itemId}", UpdateItemController()).Methods("POST")
	r.HandleFunc("/api/inventory/{itemId}", RemoveItemController()).Methods("DELETE")
}
