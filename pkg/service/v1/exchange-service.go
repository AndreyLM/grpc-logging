package v1

import (

	// "github.com/davecgh/go-spew/spew"
	"context"
	"database/sql"
	"log"
	"time"

	v1 "github.com/andreylm/grpc-logging/pkg/api/v1"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type exchangeServiceServer struct {
	db *sql.DB
}

func (s *exchangeServiceServer) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := s.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}

// NewExchangeLoggingService - creates exchange loggin service
func NewExchangeLoggingService(db *sql.DB) v1.ExchangeLogServiceServer {
	return &exchangeServiceServer{db: db}
}

func (s *exchangeServiceServer) CreateExchangeLog(ctx context.Context, req *v1.CreateExchangeLogRequest) (*v1.CreateExchangeLogResponse, error) {
	start := time.Now()
	if err := checkAPI(req.Api); err != nil {
		log.Println("Service<<CreateExchangeLog>>. Error: " + err.Error())
		return nil, err
	}
	var err error
	c, err := s.connect(ctx)
	if err != nil {
		log.Println("Service<<CreateExchangeLog>>. Error: " + err.Error())
		return nil, err
	}
	defer c.Close()

	var createdAt time.Time
	if req.ExchageLog.CreatedAt == nil {
		createdAt = time.Now()
	} else {
		createdAt, err = ptypes.Timestamp(req.ExchageLog.CreatedAt)
		if err != nil {
			log.Println("Service<<CreateExchangeLog>>. Error: " + err.Error())
			return nil, status.Error(codes.InvalidArgument, "createdAt has invalid format-> "+err.Error())
		}
	}

	_, err = c.ExecContext(ctx, "INSERT INTO exchanges(type_id, state_id, request_id, declaration_id, register_id, content, created_at)"+
		" VALUES ($1, $2, $3, $4, $5, $6, $7)",
		req.ExchageLog.TypeId, req.ExchageLog.StateId, req.ExchageLog.RequestId,
		req.ExchageLog.DeclarationId, req.ExchageLog.RegisterId, req.ExchageLog.Content, createdAt)
	if err != nil {
		log.Println("Service<<CreateExchangeLog>>. Error: " + err.Error())
		return nil, status.Error(codes.Unknown, "failed to insert-> "+err.Error())
	}

	log.Println("Service<<CreateExchangeLog>> Duration: ", time.Since(start))

	return &v1.CreateExchangeLogResponse{
		Api: apiVersion,
	}, nil
}

func (s *exchangeServiceServer) FindExchangeLogs(ctx context.Context, req *v1.FindExchangeLogsRequest) (*v1.FindExchangeLogsResponse, error) {
	start := time.Now()
	if err := checkAPI(req.Api); err != nil {
		log.Println("Service<<FindExchangeLogs>> Error: " + err.Error())
		return nil, err
	}
	c, err := s.connect(ctx)
	if err != nil {
		log.Println("Service<<FindExchangeLogs>> Error: " + err.Error())
		return nil, err
	}
	defer c.Close()

	query, err := createQuery(queryFindExchangeLog, req)
	if err != nil {
		log.Println("Service<<FindExchangeLogs>> Error: " + err.Error())
		return nil, status.Error(codes.Unknown, "failed to create query-> "+err.Error())
	}

	rows, err := c.QueryContext(ctx, query)
	if err != nil {
		log.Println("Service<<FindExchangeLogs>> Error: " + err.Error())
		return nil, status.Error(codes.Unknown, "failed to query exchange logs-> "+err.Error()+" Query: "+query)
	}
	defer rows.Close()

	var createdAt time.Time
	list := []*v1.ExchangeLog{}

	for rows.Next() {
		exchLog := new(v1.ExchangeLog)
		if err := rows.Scan(&exchLog.TypeId, &exchLog.StateId,
			&exchLog.RequestId, &exchLog.DeclarationId,
			&exchLog.RegisterId, &exchLog.Content, &createdAt); err != nil {
			log.Println("Service<<FindExchangeLogs>> Error: failed to retrieve field values-> " + err.Error())
			return nil, status.Error(codes.Unknown, "failed to retrieve field values-> "+err.Error())
		}
		exchLog.CreatedAt, err = ptypes.TimestampProto(createdAt)
		if err != nil {
			log.Println("Service<<FindExchangeLogs>> Error: createdAt field has invalid format-> " + err.Error())
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
		list = append(list, exchLog)
	}

	if err := rows.Err(); err != nil {
		log.Println("Service<<FindExchangeLogs>> Error: failed to retrieve data-> " + err.Error())
		return nil, status.Error(codes.Unknown, "failed to retrieve data-> "+err.Error())
	}

	log.Println("Service<<FindExchangeLogs>>. Duration:", time.Since(start))
	return &v1.FindExchangeLogsResponse{
		Api:         apiVersion,
		ExchageLogs: list,
	}, nil
}
