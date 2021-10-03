package main

import (
	_ "encoding/json"
	"fmt"
	_ "io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	core "github.com/nmluci/KissatenService/cafe/router"
)

func main() {
	fmt.Println("Listening at localhost:8081")
	router := mux.NewRouter()
	core.Router(router)
	log.Fatal(http.ListenAndServe(":8081", router))
}
