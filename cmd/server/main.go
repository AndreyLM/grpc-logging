package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	cmd "github.com/andreylm/grpc-logging/pkg/cmd/server"
)

func main() {
	var cfg cmd.Config
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "", "gRPC port to bind")
	flag.StringVar(&cfg.DBHost, "db-host", "", "DB host")
	flag.StringVar(&cfg.DBPort, "db-port", "", "DB port")
	flag.StringVar(&cfg.DBUser, "db-user", "", "DB user")
	flag.StringVar(&cfg.DBPassword, "db-password", "", "DB password")
	flag.StringVar(&cfg.DBSchema, "db-schema", "", "DB schema")
	flag.StringVar(&cfg.LogPath, "log-path", "log.txt", "path to log file, otherwise will be used ./log.txt")
	flag.Parse()

	logFile, err := os.OpenFile(cfg.LogPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0640)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile)

	if len(cfg.GRPCPort) == 0 {
		panic(fmt.Errorf("invalid TCP port for gRPC server: '%s'", cfg.GRPCPort))
	}

	if err := cmd.RunServer(&cfg); err != nil {
		log.Println("Error running server", err)
		os.Exit(1)
	}
}
