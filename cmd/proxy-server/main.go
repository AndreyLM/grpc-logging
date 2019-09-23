package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	cmd "github.com/andreylm/grpc-logging/pkg/cmd/proxy"
)

func main() {
	var cfg cmd.Config
	flag.StringVar(&cfg.GRPCProxyPort, "grpc-proxy-port", "", "gRPC proxy port to bind")
	flag.StringVar(&cfg.GRPCServerAddress, "grpc-server-address", "", "gRPC server port to connect")
	flag.StringVar(&cfg.LogPath, "log-path", "./log.txt", "log path, othewise will be used ./log.txt")

	flag.Parse()

	if len(cfg.GRPCProxyPort) == 0 {
		panic(fmt.Errorf("invalid TCP port for gRPC proxy server: '%s'", cfg.GRPCProxyPort))
	}

	if len(cfg.GRPCServerAddress) == 0 {
		panic(fmt.Errorf("invalid gRPC server address: '%s'", cfg.GRPCServerAddress))
	}
	logFile, err := os.OpenFile(cfg.LogPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0640)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile)

	if err := cmd.RunServer(&cfg); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
