package proxy

import (
	"context"
	"fmt"
	"time"

	v2 "github.com/andreylm/grpc-logging/pkg/api/v2"
	"github.com/andreylm/grpc-logging/pkg/request"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *loggingProxyServer) CreateExchange(ctx context.Context, req *v2.CreateExchangeRequest) (*v2.CreateExchangeResponse, error) {
	requestInfo := request.NewRequestInfo(ctx, serviceName)
	requestInfo.LogRequest()

	if err := checkAPI(req.Api); err != nil {
		requestInfo.LogError(err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := s.client.CreateExchange(ctx, req)
	if err != nil {
		requestInfo.LogError(err)
		return nil, status.Error(codes.Unknown, fmt.Sprintf("<<%s: 'CreateExchange'>> Error: %s", requestInfo.GetServiceName(), err))
	}

	requestInfo.LogDuration()
	return res, nil
}

func (s *loggingProxyServer) FindExchanges(ctx context.Context, req *v2.FindExchangesRequest) (*v2.FindExchangesResponse, error) {
	requestInfo := request.NewRequestInfo(ctx, serviceName)
	requestInfo.LogRequest()

	if err := checkAPI(req.Api); err != nil {
		requestInfo.LogError(err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := s.client.FindExchanges(ctx, req)
	if err != nil {
		requestInfo.LogError(err)
		return nil, status.Error(codes.Unknown, fmt.Sprintf("<<%s: 'FindExchanges'>> Error: %s", requestInfo.GetServiceName(), err))
	}

	requestInfo.LogDuration()
	return res, nil
}
