package network

import (
	"net/http"
)

func InternetIsConnected() bool {
	_, err := http.Get("http://clients3.google.com/generate_204")
	if err != nil {
		return false
	}
	return true
}
