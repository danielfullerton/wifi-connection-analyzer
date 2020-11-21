package main

import (
	"encoding/csv"
	"github.com/sevlyar/go-daemon"
	"log"
	"os"
	"sync"
	"time"
	"wifi-connection-analyzer/pkg/cli"
	"wifi-connection-analyzer/pkg/stringOps"
	"wifi-connection-analyzer/pkg/timer"
)

func main() {
	// get CLI arguments
	args, parseArgsErr := cli.GetCLIArgs()
	if parseArgsErr != nil {
		log.Fatal(parseArgsErr.Error())
	}

	prefixedFileLocation := stringOps.GetPrefixedFileLocation(args.FileLocation, time.Now().Format(time.RFC3339))

	// get Writer to operate on csv
	recordFile, createFileErr := os.Create(prefixedFileLocation)
	if createFileErr != nil {
		log.Fatal(createFileErr)
	}

	writer := csv.NewWriter(recordFile)

	// write headers to file
	writeHeadersErr := writer.WriteAll([][]string{
		{"Time Checked", "Connected"},
	})
	if writeHeadersErr != nil {
		log.Fatal(writeHeadersErr)
	}

	ctx := &daemon.Context{
		PidFileName: "wifiAnalyzer.pid",
		PidFilePerm: 0644,
		LogFileName: "analyzer.log",
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
	go timer.StartNetworkDaemon(args, writer)
	wg.Wait()
}
