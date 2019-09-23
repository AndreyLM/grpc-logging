package cmd

import (
	// "os"
	// "github.com/davecgh/go-spew/spew"
	// "log"
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/andreylm/grpc-logging/pkg/protocol/grpc"
	v1 "github.com/andreylm/grpc-logging/pkg/service/v1"

	// postgres dialect for sql connection
	_ "github.com/lib/pq"
)

// RunServer - runs server
func RunServer(cfg *Config) error {
	log.Println("Running server...")
	ctx := context.Background()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBSchema,
	)
	log.Println(dsn)
	db, err := sql.Open("postgres", dsn)
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(5)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Println(err)
		panic(err)
	}

	v1APIUserLoggin := v1.NewLoggingService(db)
	v1APIExchangeLoggin := v1.NewExchangeLoggingService(db)

	return grpc.RunServer(ctx, v1APIUserLoggin, v1APIExchangeLoggin, cfg.GRPCPort)
}
