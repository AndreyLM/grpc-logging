package v2

import (

	// "github.com/davecgh/go-spew/spew"
	"context"
	"database/sql"
	"errors"

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

func (s *loggingServiceServer) getCount(ctx context.Context, query string, conn *sql.Conn) (int64, error) {
	rows, err := conn.QueryContext(ctx, query)
	if err != nil {
		return 0, errors.New("error making count query " + err.Error())
	}
	defer rows.Close()
	if !rows.Next() {
		return 0, errors.New("error getting count")
	}
	var count int64
	if err := rows.Scan(&count); err != nil {
		return 0, errors.New("error getting count")
	}

	return count, nil
}

// NewLoggingService - creates loggin service
func NewLoggingService(db *sql.DB, debug bool) v2.LogginServiceServer {
	return &loggingServiceServer{
		db:    db,
		debug: debug,
	}
}
