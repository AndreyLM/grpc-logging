package main

import (
	"fmt"
	"os"

	cmd "github.com/andreylm/grpc-logging/pkg/cmd/v2/proxy"
)

func main() {
	if err := cmd.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
