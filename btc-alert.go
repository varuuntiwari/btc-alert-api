package main

import (
	"net/http"

	"github.com/gorilla/mux"
	h "github.com/varuuntiwari/btc-alert-api/handlers"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/alerts/create", h.CreateAlert).Methods("GET")
	http.Handle("/", router)

	http.ListenAndServe(":8080", router)
}
