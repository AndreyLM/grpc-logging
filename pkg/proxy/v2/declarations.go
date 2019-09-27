package proxy

import (
	"context"
	"time"

	v2 "github.com/andreylm/grpc-logging/pkg/api/v2"
	"github.com/andreylm/grpc-logging/pkg/request"
	"google.golang.org/grpc/codes"
)

func (s *loggingProxyServer) CreateDeclaration(ctx context.Context, req *v2.CreateDeclarationRequest) (*v2.CreateDeclarationResponse, error) {
	var err error
	requestInfo := request.NewRequestInfo(ctx, serviceName, "CreateDeclaration")
	if s.debug {
		requestInfo.LogRequest()
	}
	defer func() {
		if err != nil {
			requestInfo.LogError(err)
		}
		if s.debug {
			requestInfo.LogDuration()
		}
	}()

	err = checkAPI(req.Api)
	if err != nil {
		return nil, requestInfo.WrapError(codes.Unknown, err)
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	ctx = requestInfo.ContextWithMetadata(ctx)

	res, err := s.client.CreateDeclaration(ctx, req)
	if err != nil {
		return nil, requestInfo.WrapError(codes.Unknown, err)
	}
	return res, nil
}

func (s *loggingProxyServer) FindDeclarations(ctx context.Context, req *v2.FindDeclarationsRequest) (*v2.FindDeclarationsResponse, error) {
	var err error
	requestInfo := request.NewRequestInfo(ctx, serviceName, "FindDeclarations")
	if s.debug {
		requestInfo.LogRequest()
	}
	defer func() {
		if err != nil {
			requestInfo.LogError(err)
		}
		if s.debug {
			requestInfo.LogDuration()
		}
	}()

	err = checkAPI(req.Api)
	if err != nil {
		return nil, requestInfo.WrapError(codes.Unknown, err)
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	ctx = requestInfo.ContextWithMetadata(ctx)

	res, err := s.client.FindDeclarations(ctx, req)
	if err != nil {
		return nil, requestInfo.WrapError(codes.Unknown, err)
	}
	return res, nil
}
