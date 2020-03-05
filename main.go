package main

import (
	"SparePartsShop/Routers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	Routers.Routers(router)
}
