package router

import (
	"github.com/gorilla/mux"
	inv "github.com/nmluci/KissatenService/internal/inventory/controller"
	cafe "github.com/nmluci/KissatenService/internal/kissaten/controller"
	user "github.com/nmluci/KissatenService/internal/userdata/controller"
	"github.com/nmluci/KissatenService/pkg/middleware"
)

func Router() *mux.Router {
	// ctx := context.Background()
	router := mux.NewRouter()
	router.Use(middleware.NoFunc())
	cafe.RegisterKissatenSubrouter(router)
	inv.RegisterInventorySubrouter(router)
	user.RegisterUserdataSubrouter(router)
	return router
}
