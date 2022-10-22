package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	db "github.com/varuuntiwari/btc-alert-api/database"
)

type Status struct {
	Message string `json:"msg"`
}

func CreateAlert(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	err := validateURL(q)
	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Status{Message: err.Error()})
		return
	}
	priceCheck, _ := strconv.ParseInt(q.Get("priceCheck"), 10, 64)
	err = db.AddAlert(q.Get("email"), priceCheck)
	if err != nil {
		log.Print(err.Error())
		json.NewEncoder(w).Encode(Status{Message: err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Status{Message: "alert created"})
}

func DeleteAlert(w http.ResponseWriter, r *http.Request) {
	if !r.URL.Query().Has("email") {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Status{Message: "email to delete not found"})
		return
	}
	err := db.DelAlert(r.URL.Query().Get("email"))
	if err != nil {
		log.Print(err.Error())
		json.NewEncoder(w).Encode(Status{Message: err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Status{Message: "alert deleted"})
}