package database

import "log"

func AddAlert(email string, priceCheck int64) error {
	log.Printf("Adding price check of %v for email %v", priceCheck, email)
	return nil
}
