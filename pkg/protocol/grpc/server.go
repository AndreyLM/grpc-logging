package grpc

import (
	"log"
	"os"
	"context"
	"google.golang.org/grpc"
	"net"
	"github.com/andreylm/grpc-logging/pkg/api/v1"
)

// RunServer - runs server
func RunServer(ctx context.Context, v1API v1.UserLogServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	server := grpc.NewServer()
	v1.RegisterUserLogServiceServer(server, v1API)
	c := make(chan os.Signal, 1)
	go func() {
		for range c {
			log.Println("shutting down gRPC server...")
			server.GracefulStop()
			<-ctx.Done()
		}
	}()

	log.Println("starting gRPC server...")
	return server.Serve(listen)
}