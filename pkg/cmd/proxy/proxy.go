package cmd

import (
	"log"

	"context"
	"flag"
	"fmt"

	cli "github.com/andreylm/grpc-logging/pkg/api/v1"
	"github.com/andreylm/grpc-logging/pkg/protocol/grpc"
	grpc_proto "google.golang.org/grpc"

	v1 "github.com/andreylm/grpc-logging/pkg/service/v1"
)

// Config - configuration for Server
type Config struct {
	GRPCProxyPort     string
	GRPCServerAddress string
}

// RunServer - runs server
func RunServer() error {
	ctx := context.Background()
	cfg, err := getConfig()
	if err != nil {
		return err
	}

	log.Println(cfg.GRPCServerAddress)
	conn, err := grpc_proto.Dial(cfg.GRPCServerAddress, grpc_proto.WithInsecure())
	if err != nil {
		log.Println(err)
		return fmt.Errorf("<<PROXY>>: error dialing gRPC service: '%s'", err.Error())
	}
	defer conn.Close()

	cli := cli.NewUserLogServiceClient(conn)
	v1API := v1.NewLoggingProxyServer(cli)
	return grpc.RunServer(ctx, v1API, cfg.GRPCProxyPort)
}

func getConfig() (*Config, error) {
	var cfg Config
	flag.StringVar(&cfg.GRPCProxyPort, "grpc-proxy-port", "", "gRPC proxy port to bind")
	flag.StringVar(&cfg.GRPCServerAddress, "grpc-server-address", "", "gRPC server port to connect")

	flag.Parse()

	if len(cfg.GRPCProxyPort) == 0 {
		return nil, fmt.Errorf("invalid TCP port for gRPC proxy server: '%s'", cfg.GRPCProxyPort)
	}

	if len(cfg.GRPCServerAddress) == 0 {
		return nil, fmt.Errorf("invalid gRPC server address: '%s'", cfg.GRPCServerAddress)
	}

	return &cfg, nil
}
