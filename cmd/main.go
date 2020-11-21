package main

import (
	"github.com/sevlyar/go-daemon"
	"log"
	"sync"
	"wifi-connection-analyzer/pkg/cli"
	"wifi-connection-analyzer/pkg/timer"
)

func main() {
	args, parseArgsErr := cli.GetCLIArgs()
	if parseArgsErr != nil {
		log.Fatal(parseArgsErr.Error())
	}

	ctx := &daemon.Context{
		PidFileName: "wifiAnalyzer.pid",
		PidFilePerm: 0644,
		LogFileName: args.FileLocation,
		LogFilePerm: 0640,
		WorkDir: "./",
		Umask: 027,
		Args: []string{},
	}
	d, parseArgsErr := ctx.Reborn()
	if parseArgsErr != nil {
		log.Fatal("Unable to run: ", parseArgsErr)
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
