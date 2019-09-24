package proxy

import (
	"context"
	"fmt"
	"time"

	"github.com/andreylm/grpc-logging/pkg/request"

	v2 "github.com/andreylm/grpc-logging/pkg/api/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *loggingProxyServer) CreateUser(ctx context.Context, req *v2.CreateUserRequest) (*v2.CreateUserResponse, error) {
	requestInfo := request.NewRequestInfo(ctx, serviceName)
	requestInfo.LogRequest()

	if err := checkAPI(req.Api); err != nil {
		requestInfo.LogError(err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := s.client.CreateUser(ctx, req)
	if err != nil {
		requestInfo.LogError(err)
		return nil, status.Error(codes.Unknown, fmt.Sprintf("<<%s: 'CreateUser'>> Error: %s", requestInfo.GetServiceName(), err))
	}
	return res, nil
}

func (s *loggingProxyServer) FindUsers(ctx context.Context, req *v2.FindUsersRequest) (*v2.FindUsersResponse, error) {
	requestInfo := request.NewRequestInfo(ctx, serviceName)
	requestInfo.LogRequest()

	if err := checkAPI(req.Api); err != nil {
		requestInfo.LogError(err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := s.client.FindUsers(ctx, req)
	if err != nil {
		requestInfo.LogError(err)
		return nil, status.Error(codes.Unknown, fmt.Sprintf("<<%s: 'FindUsers'>> Error: %s", requestInfo.GetServiceName(), err))
	}
	requestInfo.LogDuration()
	return res, nil
}
