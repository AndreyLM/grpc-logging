package v2

import (
	"context"
	"time"

	v2 "github.com/andreylm/grpc-logging/pkg/api/v2"
	"github.com/andreylm/grpc-logging/pkg/request"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
)

func (s *loggingServiceServer) CreateExchange(ctx context.Context, req *v2.CreateExchangeRequest) (*v2.CreateExchangeResponse, error) {
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

	c, err := s.connect(ctx)
	if err != nil {
		return nil, requestInfo.WrapError(codes.Unknown, err)
	}
	defer c.Close()

	var createdAt time.Time
	if req.Exchange.CreatedAt == nil {
		createdAt = time.Now()
	} else {
		createdAt, err = ptypes.Timestamp(req.Exchange.CreatedAt)
		if err != nil {
			return nil, requestInfo.WrapError(codes.Unknown, err)
		}
	}

	exch := req.Exchange
	_, err = c.ExecContext(ctx, "INSERT INTO exchanges "+
		"(created_at, type_id, state_id, request_id, declaration_id, register_id, content)"+
		" VALUES ($1, $2, $3, $4, $5, $6, $7)",
		createdAt, exch.TypeId, exch.StateId, exch.RequestId, exch.DeclarationId, exch.RegisterId, exch.Content)
	if err != nil {
		return nil, requestInfo.WrapError(codes.Unknown, err)
	}

	return &v2.CreateExchangeResponse{
		Api:    apiVersion,
		Status: 0,
	}, nil
}

func (s *loggingServiceServer) FindExchanges(ctx context.Context, req *v2.FindExchangesRequest) (*v2.FindExchangesResponse, error) {
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

	c, err := s.connect(ctx)
	if err != nil {
		return nil, requestInfo.WrapError(codes.Unknown, err)
	}
	defer c.Close()

	query, err := createQuery(queryFindExchanges, req)
	if err != nil {
		requestInfo.LogError(err)
		return nil, requestInfo.WrapError(codes.Unknown, err)
	}

	rows, err := c.QueryContext(ctx, query)
	if err != nil {
		requestInfo.LogError(err)
		return nil, requestInfo.WrapError(codes.Unknown, err)
	}
	defer rows.Close()

	var createdAt time.Time
	list := []*v2.Exchange{}

	// SELECT created_at, type_id, state_id, request_id, declaration_id, register_id, content FROM users

	for rows.Next() {
		exch := new(v2.Exchange)
		err = rows.Scan(&createdAt, &exch.TypeId, &exch.StateId, &exch.RequestId, &exch.DeclarationId, &exch.RegisterId, &exch.Content)
		if err != nil {
			return nil, requestInfo.WrapError(codes.Unknown, err)
		}
		exch.CreatedAt, err = ptypes.TimestampProto(createdAt)
		if err != nil {
			return nil, requestInfo.WrapError(codes.Unknown, err)
		}
		list = append(list, exch)
	}

	err = rows.Err()
	if err != nil {
		return nil, requestInfo.WrapError(codes.Unknown, err)
	}

	return &v2.FindExchangesResponse{
		Api:       apiVersion,
		Exchanges: list,
	}, nil
}
