package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	h "github.com/varuuntiwari/btc-alert-api/handlers"
)

func main() {
	// p.CheckPriceAPI()
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(h.Status{Message: "this is an API for alerting about BTC prices"})
	})
	router.HandleFunc("/alerts/create", h.CreateAlert).Methods("GET")
	router.HandleFunc("/alerts/delete", h.DeleteAlert).Methods("GET")
	router.HandleFunc("/alerts/search", h.SearchAlerts).Methods("GET")
	http.Handle("/", router)

	log.Printf("running server now")
	http.ListenAndServe(":8080", router)
	log.Printf("server stopped\n")
}
