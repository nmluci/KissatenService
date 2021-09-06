package controller

import (
	"github.com/gorilla/mux"
)

func RegisterKissatenSubrouter(r *mux.Router) {
	r.HandleFunc("/api/cart/", MakeNewCartController()).Methods("POST")
	r.HandleFunc("/api/cart/{cartId}", GetCartController()).Methods("GET")
	r.HandleFunc("/api/cart/{cartId}/buy", BuyItemController()).Methods("POST")
	r.HandleFunc("/api/cart/{cartId}/return", ReturnItemController()).Methods("POST")
	r.HandleFunc("/api/cart/{cartId}/drop", DropCartController()).Methods("POST")
	r.HandleFunc("/api/cart/{cartId}/pay", PayCartController()).Methods("GET")
}
