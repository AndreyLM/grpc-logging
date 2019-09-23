package v2

import (
	"log"

	"google.golang.org/grpc/codes"
	// "github.com/davecgh/go-spew/spew"
	"context"
	"time"

	v1 "github.com/andreylm/grpc-logging/pkg/api/v1"
	"google.golang.org/grpc/status"
)

type loggingProxyServer struct {
	client v1.UserLogServiceClient
}

// NewLoggingProxyServer - creates loggin service
func NewLoggingProxyServer(client v1.UserLogServiceClient) v1.UserLogServiceServer {
	return &loggingProxyServer{client: client}
}

func (s *loggingProxyServer) CreateUserLog(ctx context.Context, req *v1.CreateUserLogRequest) (*v1.CreateUserLogResponse, error) {
	start := time.Now()
	if err := checkAPI(req.Api); err != nil {
		log.Println("Proxy<<CreateUserLog>>. Error: " + err.Error())
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := s.client.CreateUserLog(ctx, req)
	if err != nil {
		log.Println("Proxy<<CreateUserLog>>. Error: " + err.Error())
		return nil, status.Error(codes.Unknown, "<<PROXY>> error making 'CreateUserLog' request to service: "+err.Error())
	}
	log.Println("Proxy<<CreateUserLog>>. Duration: ", time.Since(start))
	return res, nil
}

func (s *loggingProxyServer) ReadUserLog(ctx context.Context, req *v1.ReadUserLogRequest) (*v1.ReadUserLogResponse, error) {
	start := time.Now()
	if err := checkAPI(req.Api); err != nil {
		log.Println("Proxy<<ReadUserLog>>. Error: " + err.Error())
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := s.client.ReadUserLog(ctx, req)
	if err != nil {
		log.Println("Proxy<<ReadUserLog>>. Error: " + err.Error())
		return nil, status.Error(codes.Unknown, "<<PROXY>> error making 'ReadUserLog' request to service: "+err.Error())
	}
	log.Println("Proxy<<ReadUserLog>>. Duration: ", time.Since(start))
	return res, nil
}

func (s *loggingProxyServer) FindUserLogs(ctx context.Context, req *v1.FindUserLogsRequest) (*v1.FindUserLogsResponse, error) {
	start := time.Now()
	if err := checkAPI(req.Api); err != nil {
		log.Println("Proxy<<FindUserLogs>>. Error: " + err.Error())
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := s.client.FindUserLogs(ctx, req)
	if err != nil {
		log.Println("Proxy<<FindUserLogs>>. Error: " + err.Error())
		return nil, status.Error(codes.Unknown, "<<PROXY>> error making 'FindUserLogs' request to service: "+err.Error())
	}

	log.Println("Proxy<<FindUserLogs>>. Duration: ", time.Since(start))
	return res, nil
}
