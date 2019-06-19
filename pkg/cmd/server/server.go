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

	"github.com/andreylm/grpc-logging/pkg/protocol/grpc"
	v1 "github.com/andreylm/grpc-logging/pkg/service/v1"

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
	log.Println("Running server...")
	ctx := context.Background()
	var cfg Config
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "", "gRPC port to bind")
	flag.StringVar(&cfg.DBHost, "db-host", "", "DB host")
	flag.StringVar(&cfg.DBPort, "db-port", "", "DB port")
	flag.StringVar(&cfg.DBUser, "db-user", "", "DB user")
	flag.StringVar(&cfg.DBPassword, "db-password", "", "DB password")
	flag.StringVar(&cfg.DBSchema, "db-schema", "", "DB schema")

	flag.Parse()

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

	v1API := v1.NewLoggingService(db)

	return grpc.RunServer(ctx, v1API, cfg.GRPCPort)
}
