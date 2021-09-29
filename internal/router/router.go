package router

import (
	"context"

	"github.com/gorilla/mux"
	inv "github.com/nmluci/KissatenService/internal/inventory/controller"
	cafe "github.com/nmluci/KissatenService/internal/kissaten/controller"
	user "github.com/nmluci/KissatenService/internal/userdata/controller"
	"github.com/nmluci/KissatenService/pkg/database"
)

func Router() *mux.Router {
	// ctx := context.Background()
	ctx := context.Background()
	router := mux.NewRouter()
	db := database.InitializeDatabase(ctx)

	router.Use()
	cafe.RegisterKissatenSubrouter(router, db)
	inv.RegisterInventorySubrouter(router, db)
	user.RegisterUserdataSubrouter(router, db)
	return router
}
