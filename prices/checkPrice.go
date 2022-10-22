package prices

import (
	"time"

	al "github.com/varuuntiwari/btc-alert-api/alert"
	db "github.com/varuuntiwari/btc-alert-api/database"
)

func CheckPriceAPI() {
	var price float64
	for {
		price = GetPrice()
		if emails := db.SearchPrice(price); emails != nil {
			for _, email := range emails {
				al.SendAlert(email)
			}
		} else {
			time.Sleep(5 * time.Second)
		}
	}
}