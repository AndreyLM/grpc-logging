package proxy

import (
	"context"
	"time"

	"github.com/andreylm/grpc-logging/pkg/request"

	v2 "github.com/andreylm/grpc-logging/pkg/api/v2"
	"google.golang.org/grpc/codes"
)

func (s *loggingProxyServer) CreateUser(ctx context.Context, req *v2.CreateUserRequest) (*v2.CreateUserResponse, error) {
	var err error
	requestInfo := request.NewRequestInfo(ctx, serviceName, "CreateUser")
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

	res, err := s.client.CreateUser(ctx, req)
	if err != nil {
		return nil, requestInfo.WrapError(codes.Unknown, err)
	}
	return res, nil
}

func (s *loggingProxyServer) FindUsers(ctx context.Context, req *v2.FindUsersRequest) (*v2.FindUsersResponse, error) {
	var err error
	requestInfo := request.NewRequestInfo(ctx, serviceName, "FindUsers")
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

	res, err := s.client.FindUsers(ctx, req)
	if err != nil {
		return nil, requestInfo.WrapError(codes.Unknown, err)
	}
	return res, nil
}
