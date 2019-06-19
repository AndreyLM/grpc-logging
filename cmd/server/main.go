package main

import (
	"log"
	"os"

	cmd "github.com/andreylm/grpc-logging/pkg/cmd/server"
)

func main() {
	logFile, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile)
	
	if err := cmd.RunServer(); err != nil {
		log.Println("Error running server", err)
		os.Exit(1)
	}
}
