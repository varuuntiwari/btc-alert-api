package main

import (
	"log"

	"github.com/varuuntiwari/btc-alert-api/prices"
)

func main() {
	currPrice := prices.GetPrice()
	log.Print(currPrice)
}
