package router

import (
	_ "context"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	// ctx := context.Background()
	router := mux.NewRouter()
	

	return router
}
