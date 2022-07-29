package handlers

import (
	"errors"
	"net/url"
)

func validateURL(u url.Values) error {
	if u.Get("email") == "" || !u.Has("email") {
		return errors.New("email not given")
	}
	if u.Get("priceCheck") == "" || !u.Has("priceCheck") {
		return errors.New("price to check not given")
	}
	return nil
}
