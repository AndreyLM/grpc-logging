package proxy

import (
	v2 "github.com/andreylm/grpc-logging/pkg/api/v2"
)

type loggingProxyServer struct {
	debug  bool
	client v2.LogginServiceClient
}

// NewLoggingProxyServer - creates loggin service
func NewLoggingProxyServer(client v2.LogginServiceClient, debug bool) v2.LogginServiceServer {
	return &loggingProxyServer{
		client: client,
		debug:  debug,
	}
}
