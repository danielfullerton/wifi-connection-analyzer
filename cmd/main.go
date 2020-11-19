package main

import (
	"github.com/sevlyar/go-daemon"
	"log"
	"sync"
	"wifi-connection-analyzer/pkg/timer"
)

func main() {
	ctx := &daemon.Context{
		PidFileName: "wifiAnalyzer.pid",
		PidFilePerm: 0644,
		LogFileName: "./NetworkAnalzyer.log",
		LogFilePerm: 0640,
		WorkDir: "./",
		Umask: 027,
		Args: []string{},
	}
	d, err := ctx.Reborn()
	if err != nil {
		log.Fatal("Unable to run: ", err)
	}
	if d != nil {
		return
	}
	defer ctx.Release()

	log.Println("-----------------")
	log.Println("daemon started")

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go timer.StartNetworkCheckTimer(20)
	wg.Wait()
}
