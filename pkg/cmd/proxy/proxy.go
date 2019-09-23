package cmd

import (
	"log"

	"context"
	"fmt"

	cli "github.com/andreylm/grpc-logging/pkg/api/v1"
	"github.com/andreylm/grpc-logging/pkg/protocol/grpc"
	grpc_proto "google.golang.org/grpc"

	v1 "github.com/andreylm/grpc-logging/pkg/service/v1"
)

// RunServer - runs server
func RunServer(cfg *Config) error {
	ctx := context.Background()

	conn, err := grpc_proto.Dial(cfg.GRPCServerAddress, grpc_proto.WithInsecure())
	if err != nil {
		log.Println(err)
		return fmt.Errorf("<<PROXY>>: error dialing gRPC service: '%s'", err.Error())
	}
	defer conn.Close()

	c := cli.NewUserLogServiceClient(conn)
	v1API := v1.NewLoggingProxyServer(c)

	c2 := cli.NewExchangeLogServiceClient(conn)
	v1API2 := v1.NewExchangeProxyServer(c2)

	return grpc.RunServer(ctx, v1API, v1API2, cfg.GRPCProxyPort)
}
