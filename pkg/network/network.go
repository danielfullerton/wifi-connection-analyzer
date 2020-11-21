package network

import (
	"net/http"
)

func InternetIsConnected(healthCheckEndpoint string) bool {
	_, err := http.Get(healthCheckEndpoint)
	if err != nil {
		return false
	}
	return true
}
