package v2

import (
	"context"

	v2 "github.com/andreylm/grpc-logging/pkg/api/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *loggingServiceServer) CreateRule(ctx context.Context, req *v2.CreateRuleRequest) (*v2.CreateRuleResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Method not implemented")
}

func (s *loggingServiceServer) FindRules(ctx context.Context, req *v2.FindRulesRequest) (*v2.FindRulesResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Method not implemented")
}

func (s *loggingServiceServer) DeleteRules(ctx context.Context, req *v2.DeleteRulesRequest) (*v2.DeleteRulesResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Method not implemented")
}
