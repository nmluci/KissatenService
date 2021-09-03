package main

import (
	_ "encoding/json"
	"fmt"
	_ "fmt"
	_ "io/ioutil"
	"log"
	"net/http"

	_ "github.com/gorilla/mux"
	"github.com/nmluci/KissatenService/pkg/router"
)

func main() {
	fmt.Print("Listening at localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", router.Router()))
}
