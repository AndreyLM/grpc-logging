package v2

import (

	// "github.com/davecgh/go-spew/spew"
	"context"
	"database/sql"

	v2 "github.com/andreylm/grpc-logging/pkg/api/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type loggingServiceServer struct {
	db    *sql.DB
	debug bool
}

func (s *loggingServiceServer) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := s.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}

// NewLoggingService - creates loggin service
func NewLoggingService(db *sql.DB, debug bool) v2.LogginServiceServer {
	return &loggingServiceServer{
		db:    db,
		debug: debug,
	}
}
