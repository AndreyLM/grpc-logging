package v1

import (

	// "github.com/davecgh/go-spew/spew"
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	v1 "github.com/andreylm/grpc-logging/pkg/api/v1"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type loggingServiceServer struct {
	db *sql.DB
}

func (s *loggingServiceServer) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := s.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}

// NewLoggingService - creates loggin service
func NewLoggingService(db *sql.DB) v1.UserLogServiceServer {
	return &loggingServiceServer{db: db}
}

func (s *loggingServiceServer) CreateUserLog(ctx context.Context, req *v1.CreateUserLogRequest) (*v1.CreateUserLogResponse, error) {
	start := time.Now()
	if err := checkAPI(req.Api); err != nil {
		log.Println("Service<<CreateUserLog>>. Error: " + err.Error())
		return nil, err
	}
	var err error
	c, err := s.connect(ctx)
	if err != nil {
		log.Println("Service<<CreateUserLog>>. Error: " + err.Error())
		return nil, err
	}
	defer c.Close()

	var createdAt time.Time
	if req.UserLog.CreatedAt == nil {
		createdAt = time.Now()
	} else {
		createdAt, err = ptypes.Timestamp(req.UserLog.CreatedAt)
		if err != nil {
			log.Println("Service<<CreateUserLog>>. Error: " + err.Error())
			return nil, status.Error(codes.InvalidArgument, "createdAt has invalid format-> "+err.Error())
		}
	}

	var id int64

	err = c.QueryRowContext(ctx, "INSERT INTO user_logs (user_id, declaration_id, type, message, created_at)"+
		" VALUES ($1, $2, $3, $4, $5) RETURNING id",
		req.UserLog.UserId, req.UserLog.DeclarationId, req.UserLog.Type, req.UserLog.UserId, createdAt).Scan(&id)
	if err != nil {
		log.Println("Service<<CreateUserLog>>. Error: " + err.Error())
		log.Println(err)
		return nil, status.Error(codes.Unknown, "failed to insert-> "+err.Error())
	}

	log.Println("Service<<CreateUserLog>> Duration: ", time.Since(start))

	return &v1.CreateUserLogResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

func (s *loggingServiceServer) ReadUserLog(ctx context.Context, req *v1.ReadUserLogRequest) (*v1.ReadUserLogResponse, error) {
	start := time.Now()
	if err := checkAPI(req.Api); err != nil {
		return nil, err
	}
	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	rows, err := c.QueryContext(ctx, "SELECT * FROM user_logs WHERE id=$1", req.Id)
	if err != nil {
		log.Println("Service<<ReadUserLog>> Error: failed to select from user_logs-> " + err.Error())
		return nil, status.Error(codes.Unknown, "failed to select from user_logs-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			log.Println("Service<<ReadUserLog>> Error: failed to retrieve data-> " + err.Error())
			return nil, status.Error(codes.Unknown, "failed to retrieve data-> "+err.Error())
		}

		log.Println(fmt.Sprintf("Service<<ReadUserLog>> Error: log with Id='%d' is not found", req.Id))
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Log with Id='%d' is not found", req.Id))
	}

	var userLog v1.UserLog
	var createdAt time.Time
	if err := rows.Scan(&userLog.Id, &userLog.UserId, &userLog.DeclarationId, &userLog.Type, &userLog.Message, &createdAt); err != nil {
		log.Println("Service<<ReadUserLog>> Error: found multiple rows with Id= " + strconv.FormatInt(userLog.Id, 10))
		return nil, status.Error(codes.Unknown, "failed to retrieve field values->"+err.Error())
	}
	userLog.CreatedAt, err = ptypes.TimestampProto(createdAt)
	if err != nil {
		log.Println("Service<<ReadUserLog>> Error: failed to retrieve data-> " + err.Error())
		return nil, status.Error(codes.Unknown, "createdAt field has invalid format->"+err.Error())
	}

	if rows.Next() {
		log.Println("Service<<ReadUserLog>> Error: found multiple rows with Id= " + strconv.FormatInt(userLog.Id, 10))
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple rows with Id='%d'", userLog.Id))
	}

	log.Println("Service<<ReadUserLog>> Duration: ", time.Since(start))
	return &v1.ReadUserLogResponse{
		Api:     apiVersion,
		UserLog: &userLog,
	}, nil
}

func (s *loggingServiceServer) FindUserLogs(ctx context.Context, req *v1.FindUserLogsRequest) (*v1.FindUserLogsResponse, error) {
	start := time.Now()
	if err := checkAPI(req.Api); err != nil {
		log.Println("Service<<FindUserLogs>> Error: " + err.Error())
		return nil, err
	}
	c, err := s.connect(ctx)
	if err != nil {
		log.Println("Service<<FindUserLogs>> Error: " + err.Error())
		return nil, err
	}
	defer c.Close()

	query, err := createQuery(queryFindUserLog, req)
	if err != nil {
		log.Println("Service<<FindUserLogs>> Error: " + err.Error())
		return nil, status.Error(codes.Unknown, "failed to create query-> "+err.Error())
	}

	rows, err := c.QueryContext(ctx, query)
	if err != nil {
		log.Println("Service<<FindUserLogs>> Error: " + err.Error())
		return nil, status.Error(codes.Unknown, "failed to select from user_logs-> "+err.Error())
	}
	defer rows.Close()

	var createdAt time.Time
	list := []*v1.UserLog{}

	for rows.Next() {
		uLog := new(v1.UserLog)
		if err := rows.Scan(&uLog.Id, &uLog.UserId, &uLog.DeclarationId, &uLog.Type, &uLog.Message, &createdAt); err != nil {
			log.Println("Service<<FindUserLogs>> Error: failed to retrieve field values-> " + err.Error())
			return nil, status.Error(codes.Unknown, "failed to retrieve field values-> "+err.Error())
		}
		uLog.CreatedAt, err = ptypes.TimestampProto(createdAt)
		if err != nil {
			log.Println("Service<<FindUserLogs>> Error: createdAt field has invalid format-> " + err.Error())
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
		list = append(list, uLog)
	}

	if err := rows.Err(); err != nil {
		log.Println("Service<<FindUserLogs>> Error: failed to retrieve data-> " + err.Error())
		return nil, status.Error(codes.Unknown, "failed to retrieve data-> "+err.Error())
	}

	log.Println("Service<<FindUserLogs>>. Duration:", time.Since(start))
	return &v1.FindUserLogsResponse{
		Api:      apiVersion,
		UserLogs: list,
	}, nil
}
