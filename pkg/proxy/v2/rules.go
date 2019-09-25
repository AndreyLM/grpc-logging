package proxy

import (
	"context"
	"time"

	v2 "github.com/andreylm/grpc-logging/pkg/api/v2"
	"github.com/andreylm/grpc-logging/pkg/request"
	"google.golang.org/grpc/codes"
)

func (s *loggingProxyServer) CreateRule(ctx context.Context, req *v2.CreateRuleRequest) (*v2.CreateRuleResponse, error) {
	var err error
	requestInfo := request.NewRequestInfo(ctx, serviceName, "CreateRule")
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

	res, err := s.client.CreateRule(ctx, req)
	if err != nil {
		return nil, requestInfo.WrapError(codes.Unknown, err)
	}

	return res, nil
}

func (s *loggingProxyServer) FindRules(ctx context.Context, req *v2.FindRulesRequest) (*v2.FindRulesResponse, error) {
	var err error
	requestInfo := request.NewRequestInfo(ctx, serviceName, "FindRules")
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

	res, err := s.client.FindRules(ctx, req)
	if err != nil {
		return nil, requestInfo.WrapError(codes.Unknown, err)
	}

	return res, nil
}
