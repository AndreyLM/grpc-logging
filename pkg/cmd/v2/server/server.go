package cmd

import (
	// "os"
	// "github.com/davecgh/go-spew/spew"
	// "log"
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/andreylm/grpc-logging/pkg/protocol/grpc/v2"
	v2 "github.com/andreylm/grpc-logging/pkg/service/v2"

	// postgres dialect for sql connection
	_ "github.com/lib/pq"
)

// Config - configuration for Server
type Config struct {
	GRPCPort   string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBSchema   string
}

// RunServer - runs server
func RunServer() error {
	var logPath string
	var cfg Config

	flag.StringVar(&logPath, "log-path", "log.txt", "DB schema")
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "", "gRPC port to bind")
	flag.StringVar(&cfg.DBHost, "db-host", "", "DB host")
	flag.StringVar(&cfg.DBPort, "db-port", "", "DB port")
	flag.StringVar(&cfg.DBUser, "db-user", "", "DB user")
	flag.StringVar(&cfg.DBPassword, "db-password", "", "DB password")
	flag.StringVar(&cfg.DBSchema, "db-schema", "", "DB schema")
	flag.Parse()

	logFile, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile)

	ctx := context.Background()

	log.Println("Running server...")

	if len(cfg.GRPCPort) == 0 {
		return fmt.Errorf("invalid TCP port for gRPC server: '%s'", cfg.GRPCPort)
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBSchema,
	)

	db, err := sql.Open("postgres", dsn)
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(5)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	v2API := v2.NewLoggingService(db)

	return grpc.RunServer(ctx, v2API, cfg.GRPCPort)
}
