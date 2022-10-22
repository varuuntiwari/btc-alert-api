package alert

import "log"

func SendAlert(email string) {
	log.Printf("%v matches price", email)
}