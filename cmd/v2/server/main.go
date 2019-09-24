package main

import (
	"log"
	"os"

	cmd "github.com/andreylm/grpc-logging/pkg/cmd/v2/server"
)

func init() {

}

func main() {

	if err := cmd.RunServer(); err != nil {
		log.Println("Error running server", err)
		os.Exit(1)
	}
}
