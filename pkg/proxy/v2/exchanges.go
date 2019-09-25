package proxy

import (
	"context"
	"time"

	v2 "github.com/andreylm/grpc-logging/pkg/api/v2"
	"github.com/andreylm/grpc-logging/pkg/request"
	"google.golang.org/grpc/codes"
)

func (s *loggingProxyServer) CreateExchange(ctx context.Context, req *v2.CreateExchangeRequest) (*v2.CreateExchangeResponse, error) {
	var err error
	requestInfo := request.NewRequestInfo(ctx, serviceName, "CreateExchange")
	requestInfo.LogRequest()
	defer func() {
		if err != nil {
			requestInfo.LogError(err)
		}
		requestInfo.LogDuration()
	}()

	err = checkAPI(req.Api)
	if err != nil {
		return nil, requestInfo.WrapError(codes.Unknown, err)
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	ctx = requestInfo.ContextWithMetadata(ctx)

	res, err := s.client.CreateExchange(ctx, req)
	if err != nil {
		return nil, requestInfo.WrapError(codes.Unknown, err)
	}

	return res, nil
}

func (s *loggingProxyServer) FindExchanges(ctx context.Context, req *v2.FindExchangesRequest) (*v2.FindExchangesResponse, error) {
	var err error
	requestInfo := request.NewRequestInfo(ctx, serviceName, "FindExchanges")
	requestInfo.LogRequest()
	defer func() {
		if err != nil {
			requestInfo.LogError(err)
		}
		requestInfo.LogDuration()
	}()

	err = checkAPI(req.Api)
	if err != nil {
		return nil, requestInfo.WrapError(codes.Unknown, err)
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	ctx = requestInfo.ContextWithMetadata(ctx)

	res, err := s.client.FindExchanges(ctx, req)
	if err != nil {
		return nil, requestInfo.WrapError(codes.Unknown, err)
	}

	requestInfo.LogDuration()
	return res, nil
}
