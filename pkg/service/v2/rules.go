package v2

import (
	"context"
	"time"

	v2 "github.com/andreylm/grpc-logging/pkg/api/v2"
	"github.com/andreylm/grpc-logging/pkg/request"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
)

func (s *loggingServiceServer) CreateRule(ctx context.Context, req *v2.CreateRuleRequest) (*v2.CreateRuleResponse, error) {
	var err error
	requestInfo := request.NewRequestInfo(ctx, serviceName, "CreateRule")
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

	c, err := s.connect(ctx)
	if err != nil {
		return nil, requestInfo.WrapError(codes.Unknown, err)
	}
	defer c.Close()

	var createdAt time.Time
	if req.Rule.CreatedAt == nil {
		createdAt = time.Now()
	} else {
		createdAt, err = ptypes.Timestamp(req.Rule.CreatedAt)
		if err != nil {
			return nil, requestInfo.WrapError(codes.Unknown, err)
		}
	}

	_, err = c.ExecContext(ctx, "INSERT INTO rules (created_at, rule_id, created_by, content, rule_number)"+
		" VALUES ($1, $2, $3, $4, $5)",
		createdAt, req.Rule.RuleId, req.Rule.CreatedBy, req.Rule.Content, req.Rule.RuleNumber)
	if err != nil {
		return nil, requestInfo.WrapError(codes.Unknown, err)
	}

	return &v2.CreateRuleResponse{
		Api:    apiVersion,
		Status: 0,
	}, nil
}

func (s *loggingServiceServer) FindRules(ctx context.Context, req *v2.FindRulesRequest) (*v2.FindRulesResponse, error) {
	var err error
	requestInfo := request.NewRequestInfo(ctx, serviceName, "FindRules")
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
	c, err := s.connect(ctx)
	if err != nil {
		return nil, requestInfo.WrapError(codes.Unknown, err)
	}
	defer c.Close()

	query, err := createQuery(queryFindRules, req)
	if err != nil {
		return nil, requestInfo.WrapError(codes.Unknown, err)
	}

	rows, err := c.QueryContext(ctx, query)
	if err != nil {
		return nil, requestInfo.WrapError(codes.Unknown, err)
	}
	defer rows.Close()

	var createdAt time.Time
	list := []*v2.Rule{}

	for rows.Next() {
		rule := new(v2.Rule)
		err = rows.Scan(&createdAt, &rule.RuleId, &rule.CreatedBy, &rule.Content, &rule.RuleNumber)
		if err != nil {
			return nil, requestInfo.WrapError(codes.Unknown, err)
		}
		rule.CreatedAt, err = ptypes.TimestampProto(createdAt)
		if err != nil {
			return nil, requestInfo.WrapError(codes.Unknown, err)
		}
		list = append(list, rule)
	}

	err = rows.Err()
	if err != nil {
		return nil, requestInfo.WrapError(codes.Unknown, err)
	}

	response := &v2.FindRulesResponse{
		Api:   apiVersion,
		Rules: list,
	}

	query, err = createQuery(queryTotalCountRules, req)
	if err != nil {
		return response, nil
	}

	response.TotalCount, err = s.getCount(ctx, query, c)
	return response, nil
}
