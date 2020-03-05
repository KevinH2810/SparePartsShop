package Routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Routers(routes *mux.Router) {
	routes.HandleFunc("/Item", GetItems).Methods("GET")
	routes.HandleFunc("/Item/{id}", GetItem).Methods("GET")
	routes.HandleFunc("/Item", CreateItem).Methods("POST")
	routes.HandleFunc("/Item", UpdateItem).Methods("PUT")
	routes.HandleFunc("/Item", DeleteItem).Methods("DELETE")

	http.ListenAndServe(":9090", routes)
}
