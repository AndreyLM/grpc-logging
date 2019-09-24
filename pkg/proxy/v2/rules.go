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

func (s *loggingProxyServer) CreateRule(ctx context.Context, req *v2.CreateRuleRequest) (*v2.CreateRuleResponse, error) {
	requestInfo := request.NewRequestInfo(ctx, serviceName)
	requestInfo.LogRequest()

	if err := checkAPI(req.Api); err != nil {
		requestInfo.LogError(err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := s.client.CreateRule(ctx, req)
	if err != nil {
		requestInfo.LogError(err)
		return nil, status.Error(codes.Unknown, fmt.Sprintf("<<%s: 'CreateRule'>> Error: %s", requestInfo.GetServiceName(), err))
	}

	requestInfo.LogDuration()
	return res, nil
}

func (s *loggingProxyServer) FindRules(ctx context.Context, req *v2.FindRulesRequest) (*v2.FindRulesResponse, error) {
	requestInfo := request.NewRequestInfo(ctx, serviceName)
	requestInfo.LogRequest()

	if err := checkAPI(req.Api); err != nil {
		requestInfo.LogError(err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := s.client.FindRules(ctx, req)
	if err != nil {
		requestInfo.LogError(err)
		return nil, status.Error(codes.Unknown, fmt.Sprintf("<<%s: 'FindRules'>> Error: %s", requestInfo.GetServiceName(), err))
	}

	requestInfo.LogDuration()
	return res, nil
}
