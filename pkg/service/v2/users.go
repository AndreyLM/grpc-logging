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

func (s *loggingServiceServer) CreateUser(ctx context.Context, req *v2.CreateUserRequest) (*v2.CreateUserResponse, error) {
	start := time.Now()
	if err := checkAPI(req.Api); err != nil {
		log.Println("Service<<CreateUser>>. Error: " + err.Error())
		return nil, err
	}
	var err error
	c, err := s.connect(ctx)
	if err != nil {
		log.Println("Service<<CreateUser>>. Error: " + err.Error())
		return nil, err
	}
	defer c.Close()

	var createdAt time.Time
	if req.User.CreatedAt == nil {
		createdAt = time.Now()
	} else {
		createdAt, err = ptypes.Timestamp(req.User.CreatedAt)
		if err != nil {
			log.Println("Service<<CreateUser>>. Error: " + err.Error())
			return nil, status.Error(codes.InvalidArgument, "createdAt has invalid format-> "+err.Error())
		}
	}

	_, err = c.ExecContext(ctx, "INSERT INTO users (created_at, user_id, type_id, content)"+
		" VALUES ($1, $2, $3, $4)",
		createdAt, req.User.UserId, req.User.TypeId, req.User.Content)
	if err != nil {
		log.Println("Service<<CreateUser>>. Error: " + err.Error())
		return nil, status.Error(codes.Unknown, "failed to insert-> "+err.Error())
	}

	log.Println("Service<<CreateUser>> Duration: ", time.Since(start))

	return &v2.CreateUserResponse{
		Api:    apiVersion,
		Status: 0,
	}, nil
}

func (s *loggingServiceServer) FindUsers(ctx context.Context, req *v2.FindUsersRequest) (*v2.FindUsersResponse, error) {
	start := time.Now()
	if err := checkAPI(req.Api); err != nil {
		log.Println("Service<<FindUsers>> Error: " + err.Error())
		return nil, err
	}
	c, err := s.connect(ctx)
	if err != nil {
		log.Println("Service<<FindUsers>> Error: " + err.Error())
		return nil, err
	}
	defer c.Close()

	query, err := createQuery(queryFindUsers, req)
	if err != nil {
		log.Println("Service<<FindUsers>> Error: " + err.Error())
		return nil, status.Error(codes.Unknown, "failed to create query-> "+err.Error())
	}

	rows, err := c.QueryContext(ctx, query)
	if err != nil {
		log.Println("Service<<FindUsers>> Error: " + err.Error())
		return nil, status.Error(codes.Unknown, "failed to select from users-> "+err.Error())
	}
	defer rows.Close()

	var createdAt time.Time
	list := []*v2.User{}

	for rows.Next() {
		uLog := new(v2.User)
		if err := rows.Scan(&createdAt, &uLog.UserId, &uLog.TypeId, &uLog.Content); err != nil {
			log.Println("Service<<FindUsers>> Error: failed to retrieve field values-> " + err.Error())
			return nil, status.Error(codes.Unknown, "failed to retrieve field values-> "+err.Error())
		}
		uLog.CreatedAt, err = ptypes.TimestampProto(createdAt)
		if err != nil {
			log.Println("Service<<FindUsers>> Error: createdAt field has invalid format-> " + err.Error())
			return nil, status.Error(codes.Unknown, "createdAt field has invalid format-> "+err.Error())
		}
		list = append(list, uLog)
	}

	if err := rows.Err(); err != nil {
		log.Println("Service<<FindUsers>> Error: failed to retrieve data-> " + err.Error())
		return nil, status.Error(codes.Unknown, "failed to retrieve data-> "+err.Error())
	}

	log.Println("Service<<FindUsers>>. Duration:", time.Since(start))
	return &v2.FindUsersResponse{
		Api:   apiVersion,
		Users: list,
	}, nil
}
