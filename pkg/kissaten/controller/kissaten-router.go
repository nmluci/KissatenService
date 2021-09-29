package controller

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/nmluci/KissatenService/internal/kissaten/models"
)

func RegisterKissatenSubrouter(r *mux.Router, db *sql.DB) {
	km := &models.KissatenModel{DB: db}
	r.HandleFunc("/api/kissaten/make", MakeNewCartController(km)).Methods("POST")
	r.HandleFunc("/api/kissaten/cart", GetAllCartController(km)).Methods("POST")
	r.HandleFunc("/api/kissaten/cart/{cartId}", GetCartController(km)).Methods("GET")
	r.HandleFunc("/api/kissaten/cart/{cartId}/buy", BuyItemController(km)).Methods("POST")
	r.HandleFunc("/api/kissaten/cart/{cartId}/return", ReturnItemController(km)).Methods("POST")
	r.HandleFunc("/api/kissaten/cart/{cartId}/drop", DropCartController(km)).Methods("POST")
	r.HandleFunc("/api/kissaten/cart/{cartId}/pay", PayCartController(km)).Methods("GET")
}
