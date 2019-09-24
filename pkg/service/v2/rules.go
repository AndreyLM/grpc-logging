package v2

import (
	"context"
	"log"
	"time"

	v2 "github.com/andreylm/grpc-logging/pkg/api/v2"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *loggingServiceServer) CreateRule(ctx context.Context, req *v2.CreateRuleRequest) (*v2.CreateRuleResponse, error) {
	start := time.Now()
	if err := checkAPI(req.Api); err != nil {
		log.Println("Service<<CreateRule>>. Error: " + err.Error())
		return nil, err
	}
	var err error
	c, err := s.connect(ctx)
	if err != nil {
		log.Println("Service<<CreateRule>>. Error: " + err.Error())
		return nil, err
	}
	defer c.Close()

	var createdAt time.Time
	if req.Rule.CreatedAt == nil {
		createdAt = time.Now()
	} else {
		createdAt, err = ptypes.Timestamp(req.Rule.CreatedAt)
		if err != nil {
			log.Println("Service<<CreateRule>>. Error: " + err.Error())
			return nil, status.Error(codes.InvalidArgument, "createdAt has invalid format-> "+err.Error())
		}
	}

	_, err = c.ExecContext(ctx, "INSERT INTO rules (created_at, rule_id, created_by, content, rule_number)"+
		" VALUES ($1, $2, $3, $4, $5)",
		createdAt, req.Rule.RuleId, req.Rule.CreatedBy, req.Rule.Content, req.Rule.RuleNumber)
	if err != nil {
		log.Println("Service<<CreateRule>>. Error: " + err.Error())
		return nil, status.Error(codes.Unknown, "failed to insert-> "+err.Error())
	}

	log.Println("Service<<CreateRule>> Duration: ", time.Since(start))

	return &v2.CreateRuleResponse{
		Api:    apiVersion,
		Status: 0,
	}, nil
}

func (s *loggingServiceServer) FindRules(ctx context.Context, req *v2.FindRulesRequest) (*v2.FindRulesResponse, error) {
	start := time.Now()
	if err := checkAPI(req.Api); err != nil {
		log.Println("Service<<FindRules>> Error: " + err.Error())
		return nil, err
	}
	c, err := s.connect(ctx)
	if err != nil {
		log.Println("Service<<FindRules>> Error: " + err.Error())
		return nil, err
	}
	defer c.Close()

	query, err := createQuery(queryFindRules, req)
	if err != nil {
		log.Println("Service<<FindRules>> Error: " + err.Error())
		return nil, status.Error(codes.Unknown, "failed to create query-> "+err.Error())
	}

	rows, err := c.QueryContext(ctx, query)
	if err != nil {
		log.Println("Service<<FindRules>> Error: " + err.Error())
		return nil, status.Error(codes.Unknown, "failed to select from rules-> "+err.Error())
	}
	defer rows.Close()

	var createdAt time.Time
	list := []*v2.Rule{}

	for rows.Next() {
		rule := new(v2.Rule)
		if err := rows.Scan(&createdAt, &rule.RuleId, &rule.CreatedBy, &rule.Content, &rule.RuleNumber); err != nil {
			log.Println("Service<<FindRules>> Error: failed to retrieve field values-> " + err.Error())
			return nil, status.Error(codes.Unknown, "failed to retrieve field values-> "+err.Error())
		}
		rule.CreatedAt, err = ptypes.TimestampProto(createdAt)
		if err != nil {
			log.Println("Service<<FindRules>> Error: createdAt field has invalid format-> " + err.Error())
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
		list = append(list, rule)
	}

	if err := rows.Err(); err != nil {
		log.Println("Service<<FindRules>> Error: failed to retrieve data-> " + err.Error())
		return nil, status.Error(codes.Unknown, "failed to retrieve data-> "+err.Error())
	}

	log.Println("Service<<FindRules>>. Duration:", time.Since(start))
	return &v2.FindRulesResponse{
		Api:   apiVersion,
		Rules: list,
	}, nil
}
