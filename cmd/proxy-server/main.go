package main

import (
	"fmt"
	"log"
	"os"

	cmd "github.com/andreylm/grpc-logging/pkg/cmd/proxy"
)

func main() {
	logFile, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile)

	if err := cmd.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
