package main

import (
	"net/http"

	"SparePartsShop/Routers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	Routers.Routes(router)

	http.ListenAndServe(":9090", router)
}
