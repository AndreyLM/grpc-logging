package main

import (
	"os"
	"fmt"
	"github.com/andreylm/grpc-logging/pkg/cmd/server"
)

func main() {
	if err := cmd.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}