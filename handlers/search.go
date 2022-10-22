package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	db "github.com/varuuntiwari/btc-alert-api/database"
)

func SearchAlerts(w http.ResponseWriter, r *http.Request) {
	if !r.URL.Query().Has("email") {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(Status{Message: "email to delete not found"})
		return
	}
	emailSearch := r.URL.Query().Get("email")

	res, err := db.SearchAlerts(emailSearch)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(Status{Message: err.Error()})
		return
	} else {
		log.Println(res)
	}
}
