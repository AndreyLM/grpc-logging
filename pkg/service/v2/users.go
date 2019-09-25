package v2

import (
	"context"
	"time"

	v2 "github.com/andreylm/grpc-logging/pkg/api/v2"
	"github.com/andreylm/grpc-logging/pkg/request"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
)

func (s *loggingServiceServer) CreateUser(ctx context.Context, req *v2.CreateUserRequest) (*v2.CreateUserResponse, error) {
	var err error
	requestInfo := request.NewRequestInfo(ctx, serviceName, "CreateUser")
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
	if req.User.CreatedAt == nil {
		createdAt = time.Now()
	} else {
		createdAt, err = ptypes.Timestamp(req.User.CreatedAt)
		if err != nil {
			return nil, requestInfo.WrapError(codes.Unknown, err)
		}
	}

	_, err = c.ExecContext(ctx, "INSERT INTO users (created_at, user_id, type_id, content)"+
		" VALUES ($1, $2, $3, $4)",
		createdAt, req.User.UserId, req.User.TypeId, req.User.Content)
	if err != nil {
		return nil, requestInfo.WrapError(codes.Unknown, err)
	}

	return &v2.CreateUserResponse{
		Api:    apiVersion,
		Status: 0,
	}, nil
}

func (s *loggingServiceServer) FindUsers(ctx context.Context, req *v2.FindUsersRequest) (*v2.FindUsersResponse, error) {
	var err error
	requestInfo := request.NewRequestInfo(ctx, serviceName, "FindUsers")
	requestInfo.LogRequest()
	defer func() {
		if err != nil {
			requestInfo.LogError(err)
		}
		requestInfo.LogDuration()
	}()

	err = checkAPI(req.Api)
	if err != nil {
		return nil, err
	}

	c, err := s.connect(ctx)
	if err != nil {
		return nil, requestInfo.WrapError(codes.Unknown, err)
	}
	defer c.Close()

	query, err := createQuery(queryFindUsers, req)
	if err != nil {
		return nil, requestInfo.WrapError(codes.Unknown, err)
	}

	rows, err := c.QueryContext(ctx, query)
	if err != nil {
		return nil, requestInfo.WrapError(codes.Unknown, err)
	}
	defer rows.Close()

	var createdAt time.Time
	list := []*v2.User{}

	for rows.Next() {
		uLog := new(v2.User)
		err = rows.Scan(&createdAt, &uLog.UserId, &uLog.TypeId, &uLog.Content)
		if err != nil {
			return nil, requestInfo.WrapError(codes.Unknown, err)
		}
		uLog.CreatedAt, err = ptypes.TimestampProto(createdAt)
		if err != nil {
			return nil, requestInfo.WrapError(codes.Unknown, err)
		}
		list = append(list, uLog)
	}

	err = rows.Err()
	if err != nil {
		return nil, requestInfo.WrapError(codes.Unknown, err)
	}

	return &v2.FindUsersResponse{
		Api:   apiVersion,
		Users: list,
	}, nil
}
