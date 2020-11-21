package timer

import (
	"encoding/csv"
	"log"
	"time"
	"wifi-connection-analyzer/pkg/network"
	"wifi-connection-analyzer/pkg/types"
)

var count = 0

func networkJob(healthCheckEndpoint string) string {
	count++
	isConnected := network.InternetIsConnected(healthCheckEndpoint)
	var msg string
	if isConnected {
		msg = "Network is connected"
	} else {
		msg = "Network is NOT connected"
	}
	return msg
}

func StartNetworkDaemon(options types.Options, writer *csv.Writer) {
	ticker := time.NewTicker(time.Duration(options.IntervalSeconds) * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
				case <- ticker.C: {
					msg := networkJob(options.HealthCheckEndpoint)
					_ = writer.WriteAll([][]string{
						{time.Now().Format(time.RFC1123), msg},
					})
					log.Println("networkJob ran successfully")
				}
				case <- quit: {
					ticker.Stop()
					return
				}
			}
		}
	}()
}
