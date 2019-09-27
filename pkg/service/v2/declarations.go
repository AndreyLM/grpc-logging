package v2

import (
	"context"
	"time"

	v2 "github.com/andreylm/grpc-logging/pkg/api/v2"
	"github.com/andreylm/grpc-logging/pkg/request"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
)

func (s *loggingServiceServer) CreateDeclaration(ctx context.Context, req *v2.CreateDeclarationRequest) (*v2.CreateDeclarationResponse, error) {
	var err error
	requestInfo := request.NewRequestInfo(ctx, serviceName, "CreateDeclaration")
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
	if req.Declaration.CreatedAt == nil {
		createdAt = time.Now()
	} else {
		createdAt, err = ptypes.Timestamp(req.Declaration.CreatedAt)
		if err != nil {
			return nil, requestInfo.WrapError(codes.Unknown, err)
		}
	}

	_, err = c.ExecContext(ctx, "INSERT INTO declarations (created_at, declaration_id, content, user_id, user_ip)"+
		" VALUES ($1, $2, $3, $4, $5)",
		createdAt, req.Declaration.DeclarationId, req.Declaration.Content, req.Declaration.UserId, req.Declaration.UserIp)
	if err != nil {
		return nil, requestInfo.WrapError(codes.Unknown, err)
	}

	return &v2.CreateDeclarationResponse{
		Api:    apiVersion,
		Status: 0,
	}, nil
}

func (s *loggingServiceServer) FindDeclarations(ctx context.Context, req *v2.FindDeclarationsRequest) (*v2.FindDeclarationsResponse, error) {
	var err error
	requestInfo := request.NewRequestInfo(ctx, serviceName, "FindDeclarations")
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

	query, err := createQuery(queryFindDeclarations, req)
	if err != nil {
		return nil, requestInfo.WrapError(codes.Unknown, err)
	}

	rows, err := c.QueryContext(ctx, query)
	if err != nil {
		return nil, requestInfo.WrapError(codes.Unknown, err)
	}
	defer rows.Close()

	var createdAt time.Time
	list := []*v2.Declaration{}

	for rows.Next() {
		decl := new(v2.Declaration)
		err = rows.Scan(&createdAt, &decl.DeclarationId, &decl.Content, &decl.UserId, &decl.UserIp)
		if err != nil {
			return nil, requestInfo.WrapError(codes.Unknown, err)
		}
		decl.CreatedAt, err = ptypes.TimestampProto(createdAt)
		if err != nil {
			return nil, requestInfo.WrapError(codes.Unknown, err)
		}
		list = append(list, decl)
	}

	err = rows.Err()
	if err != nil {
		return nil, requestInfo.WrapError(codes.Unknown, err)
	}

	return &v2.FindDeclarationsResponse{
		Api:          apiVersion,
		Declarations: list,
	}, nil
}
