package router

import (
	"log"

	"github.com/gorilla/mux"
	controller "github.com/nmluci/KissatenService/cafe/controllers"
)

func Router(r *mux.Router) {
	if err := controller.RegisterService(); err != nil {
		log.Fatal(err)
	}

	r.HandleFunc("/api/kissaten/make", controller.MakeNewCart()).Methods("POST")
	r.HandleFunc("/api/kissaten/cart", controller.GetAllCart()).Methods("POST")
	r.HandleFunc("/api/kissaten/cart/{cartId}", controller.GetCart()).Methods("GET")
	r.HandleFunc("/api/kissaten/cart/{cartId}/buy", controller.BuyItem()).Methods("POST")
	r.HandleFunc("/api/kissaten/cart/{cartId}/return", controller.ReturnItem()).Methods("POST")
	r.HandleFunc("/api/kissaten/cart/{cartId}/drop", controller.DropCart()).Methods("POST")
	r.HandleFunc("/api/kissaten/cart/{cartId}/pay", controller.PayCart()).Methods("GET")
}
