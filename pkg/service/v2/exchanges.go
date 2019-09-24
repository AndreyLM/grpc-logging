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

func (s *loggingServiceServer) CreateExchange(ctx context.Context, req *v2.CreateExchangeRequest) (*v2.CreateExchangeResponse, error) {
	start := time.Now()
	if err := checkAPI(req.Api); err != nil {
		log.Println("Service<<CreateExchange>>. Error: " + err.Error())
		return nil, err
	}
	var err error
	c, err := s.connect(ctx)
	if err != nil {
		log.Println("Service<<CreateExchange>>. Error: " + err.Error())
		return nil, err
	}
	defer c.Close()

	var createdAt time.Time
	if req.Exchange.CreatedAt == nil {
		createdAt = time.Now()
	} else {
		createdAt, err = ptypes.Timestamp(req.Exchange.CreatedAt)
		if err != nil {
			log.Println("Service<<CreateExchange>>. Error: " + err.Error())
			return nil, status.Error(codes.InvalidArgument, "createdAt has invalid format-> "+err.Error())
		}
	}

	exch := req.Exchange
	_, err = c.ExecContext(ctx, "INSERT INTO exchanges "+
		"(created_at, type_id, state_id, request_id, declaration_id, register_id, content)"+
		" VALUES ($1, $2, $3, $4, $5, $6, $7)",
		createdAt, exch.TypeId, exch.StateId, exch.RequestId, exch.DeclarationId, exch.RegisterId, exch.Content)
	if err != nil {
		log.Println("Service<<CreateExchange>>. Error: " + err.Error())
		return nil, status.Error(codes.Unknown, "failed to insert-> "+err.Error())
	}

	log.Println("Service<<CreateExchange>> Duration: ", time.Since(start))

	return &v2.CreateExchangeResponse{
		Api:    apiVersion,
		Status: 0,
	}, nil
}

func (s *loggingServiceServer) FindExchanges(ctx context.Context, req *v2.FindExchangesRequest) (*v2.FindExchangesResponse, error) {
	start := time.Now()
	if err := checkAPI(req.Api); err != nil {
		log.Println("Service<<FindExchanges>> Error: " + err.Error())
		return nil, err
	}
	c, err := s.connect(ctx)
	if err != nil {
		log.Println("Service<<FindExchanges>> Error: " + err.Error())
		return nil, err
	}
	defer c.Close()

	query, err := createQuery(queryFindExchanges, req)
	if err != nil {
		log.Println("Service<<FindExchanges>> Error: " + err.Error())
		return nil, status.Error(codes.Unknown, "failed to create query-> "+err.Error())
	}

	rows, err := c.QueryContext(ctx, query)
	if err != nil {
		log.Println("Service<<FindExchanges>> Error: " + err.Error())
		return nil, status.Error(codes.Unknown, "failed to select from exchanges-> "+err.Error())
	}
	defer rows.Close()

	var createdAt time.Time
	list := []*v2.Exchange{}

	// SELECT created_at, type_id, state_id, request_id, declaration_id, register_id, content FROM users

	for rows.Next() {
		exch := new(v2.Exchange)
		if err := rows.Scan(&createdAt, &exch.TypeId, &exch.StateId, &exch.RequestId, &exch.DeclarationId, &exch.RegisterId, &exch.Content); err != nil {
			log.Println("Service<<FindExchanges>> Error: failed to retrieve field values-> " + err.Error())
			return nil, status.Error(codes.Unknown, "failed to retrieve field values-> "+err.Error())
		}
		exch.CreatedAt, err = ptypes.TimestampProto(createdAt)
		if err != nil {
			log.Println("Service<<FindExchanges>> Error: createdAt field has invalid format-> " + err.Error())
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
		list = append(list, exch)
	}

	if err := rows.Err(); err != nil {
		log.Println("Service<<FindExchanges>> Error: failed to retrieve data-> " + err.Error())
		return nil, status.Error(codes.Unknown, "failed to retrieve data-> "+err.Error())
	}

	log.Println("Service<<FindExchanges>>. Duration:", time.Since(start))
	return &v2.FindExchangesResponse{
		Api:       apiVersion,
		Exchanges: list,
	}, nil
}
