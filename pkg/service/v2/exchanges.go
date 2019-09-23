package v2

import (
	"context"

	v2 "github.com/andreylm/grpc-logging/pkg/api/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *loggingServiceServer) CreateExchange(ctx context.Context, req *v2.CreateExchangeRequest) (*v2.CreateExchangeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Method not implemented")
}

func (s *loggingServiceServer) FindExchanges(ctx context.Context, req *v2.FindExchangesRequest) (*v2.FindExchangesResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Method not implemented")
}

func (s *loggingServiceServer) DeleteExchanges(ctx context.Context, req *v2.DeleteExchangesRequest) (*v2.DeleteExchangesResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Method not implemented")
}
