package main

import (
	"github.com/golang/protobuf/ptypes"
	"time"
	"context"
	"github.com/andreylm/grpc-logging/pkg/api/v1"
	"log"
	"flag"
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

	// Call Create
	req1 := v1.CreateUserLogRequest{
		Api: apiVersion,
		UserLog: &v1.UserLog{
			UserId: 1,
			DeclarationId: 4,
			Type: "Deleting some data",
			CreatedAt: createdAt,
			Message: "Testing grpc service",
		},
	}
	res1, err := cli.CreateUserLog(ctx, &req1)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create result: <%+v>\n\n", res1)

	id := res1.Id

	// Read
	req2 := v1.ReadUserLogRequest{
		Api: apiVersion,
		Id:  id,
	}
	res2, err := cli.ReadUserLog(ctx, &req2)
	if err != nil {
		log.Fatalf("Read failed: %v", err)
	}
	log.Println(res2)

	// FIND
	timeFrom, _ := ptypes.TimestampProto(time.Now().Add(-time.Minute*20))
	req3 := &v1.FindUserLogsRequest{
		Api: apiVersion,
		From: timeFrom,
	}

	res4, err := cli.FindUserLogs(ctx, req3)
	if err != nil {
		log.Fatalf("Find failed: %v", err)
	}
	log.Printf("Find result: <%+v>\n\n", res4)
}