package v1

import (
	"log"

	"google.golang.org/grpc/codes"
	// "github.com/davecgh/go-spew/spew"
	"context"
	"time"

	v1 "github.com/andreylm/grpc-logging/pkg/api/v1"
	"google.golang.org/grpc/status"
)

type exchangeProxyServer struct {
	client v1.ExchangeLogServiceClient
}

// NewExchangeProxyServer - creates loggin service
func NewExchangeProxyServer(client v1.ExchangeLogServiceClient) v1.ExchangeLogServiceServer {
	return &exchangeProxyServer{client: client}
}

func (s *exchangeProxyServer) CreateExchangeLog(ctx context.Context, req *v1.CreateExchangeLogRequest) (*v1.CreateExchangeLogResponse, error) {
	start := time.Now()
	if err := checkAPI(req.Api); err != nil {
		log.Println("Proxy<<CreateExchangeLog>>. Error: " + err.Error())
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := s.client.CreateExchangeLog(ctx, req)
	if err != nil {
		log.Println("Proxy<<CreateExchangeLog>>. Error: " + err.Error())
		return nil, status.Error(codes.Unknown, "<<PROXY>> error making 'CreateExchangeLog' request to service: "+err.Error())
	}
	log.Println("Proxy<<CreateExchangeLog>>. Duration: ", time.Since(start))
	return res, nil
}

func (s *exchangeProxyServer) FindExchangeLogs(ctx context.Context, req *v1.FindExchangeLogsRequest) (*v1.FindExchangeLogsResponse, error) {
	start := time.Now()
	if err := checkAPI(req.Api); err != nil {
		log.Println("Proxy<<FindExchangeLogs>>. Error: " + err.Error())
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := s.client.FindExchangeLogs(ctx, req)
	if err != nil {
		log.Println("Proxy<<FindExchangeLogs>>. Error: " + err.Error())
		return nil, status.Error(codes.Unknown, "<<PROXY>> error making 'FindExchangeLogs' request to service: "+err.Error())
	}

	log.Println("Proxy<<FindExchangeLogs>>. Duration: ", time.Since(start))
	return res, nil
}
