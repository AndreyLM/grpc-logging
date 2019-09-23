package main

import (
	"context"
	"flag"
	"log"
	"time"

	v1 "github.com/andreylm/grpc-logging/pkg/api/v1"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
)

const (
	apiVersion = "v1"
)

func main() {
	// get configuration
	address := flag.String("server", "", "gRPC server in format host:port")
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	cli := v1.NewUserLogServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	createdAt, _ := ptypes.TimestampProto(time.Now())

	req1 := v1.CreateUserLogRequest{
		Api: apiVersion,
		UserLog: &v1.UserLog{
			UserId:        1,
			DeclarationId: 4,
			Type:          "Deleting some data",
			CreatedAt:     createdAt,
			Message:       "Testing grpc service",
		},
	}
	res1, err := cli.CreateUserLog(ctx, &req1)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create result: <%+v>\n\n", res1)

	req2 := v1.ReadUserLogRequest{
		Api: apiVersion,
		Id:  1,
	}
	res2, err := cli.ReadUserLog(ctx, &req2)
	if err != nil {
		log.Fatalf("Read failed: %v", err)
	}
	log.Println(res2)

	// FIND
	timeFrom, _ := ptypes.TimestampProto(time.Now().Add(-time.Minute * 20))
	req3 := &v1.FindUserLogsRequest{
		Api:  apiVersion,
		From: timeFrom,
	}

	res4, err := cli.FindUserLogs(ctx, req3)
	if err != nil {
		log.Fatalf("Find failed: %v", err)
	}
	log.Printf("Find result: <%+v>\n\n", res4)

	cli2 := v1.NewExchangeLogServiceClient(conn)
	req4 := v1.CreateExchangeLogRequest{
		Api: apiVersion,
		ExchageLog: &v1.ExchangeLog{
			TypeId:        3,
			StateId:       3,
			RequestId:     3,
			DeclarationId: 3,
			RegisterId:    "Register ID",
			Content:       "Some content",
		},
	}
	res5, err := cli2.CreateExchangeLog(ctx, &req4)
	if err != nil {
		log.Fatalf("Create Failed: %v", err)
	}
	log.Printf("Create response: <%+v>\n\n", res5)

	req5 := v1.FindExchangeLogsRequest{
		Api:    apiVersion,
		TypeId: 3,
	}
	res6, err := cli2.FindExchangeLogs(ctx, &req5)
	if err != nil {
		log.Fatalf("<<Exchange service>> Error find: <%+v>", err)
	}
	log.Printf("<<Exchange service>> Find respose: <%+v>", res6)
}
