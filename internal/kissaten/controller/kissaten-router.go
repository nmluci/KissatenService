package controller

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/nmluci/KissatenService/internal/kissaten/models"
)

func RegisterKissatenSubrouter(r *mux.Router, db *sql.DB) {

	km := &models.KissatenModel{DB: db}
	r.HandleFunc("/api/cart/", MakeNewCartController(km)).Methods("POST")
	r.HandleFunc("/api/cart/{cartId}", GetCartController(km)).Methods("GET")
	r.HandleFunc("/api/cart/{cartId}/buy", BuyItemController(km)).Methods("POST")
	r.HandleFunc("/api/cart/{cartId}/return", ReturnItemController(km)).Methods("POST")
	r.HandleFunc("/api/cart/{cartId}/drop", DropCartController(km)).Methods("POST")
	r.HandleFunc("/api/cart/{cartId}/pay", PayCartController(km)).Methods("GET")
}
