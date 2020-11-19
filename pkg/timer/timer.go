package timer

import (
	"log"
	"time"
	"wifi-connection-analyzer/pkg/network"
)

var count = 0

func networkJob() {
	count++
	isConnected := network.InternetIsConnected()
	var msg string
	if isConnected {
		msg = "Network is connected"
	} else {
		msg = "Network is NOT connected"
	}
	log.Println(msg)
}

func StartNetworkCheckTimer(intervalSecs int) {
	ticker := time.NewTicker(time.Duration(intervalSecs) * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
				case <- ticker.C: {
					networkJob()
				}
				case <- quit: {
					ticker.Stop()
					return
				}
			}
		}
	}()
}
