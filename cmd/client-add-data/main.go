package main

import (
	"strconv"
	"sync"
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	var errors, success int64
	tChan := time.Tick(5*time.Second)
	go func(c <-chan time.Time){
		for range c {
			log.Printf("Errors: %d\n", errors)
			log.Printf("Successes: %d\n", success)
		}
	}(tChan)
	
	var mu sync.Mutex
	var wg sync.WaitGroup
	start := time.Now()
	sema := make(chan struct{}, 100)
	
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func(i int, se chan struct{}, wg *sync.WaitGroup) {
			se <-struct{}{}
				
			createdAt, _ := ptypes.TimestampProto(time.Now())

			// Call Create
			req1 := v1.CreateUserLogRequest{
				Api: apiVersion,
				UserLog: &v1.UserLog{
					UserId: int64(i),
					DeclarationId: int64(i),
					Type: "Deleting some data " + strconv.Itoa(i),
					CreatedAt: createdAt,
					Message: "Testing grpc service",
				},
			}
		
			_, err := cli.CreateUserLog(ctx, &req1)
			mu.Lock()
			if err != nil {
				errors++
				log.Panicln(err)
			} else {
				success++
			}
			mu.Unlock()
			wg.Done()
			<- se
		}(i, sema, &wg)
		
	}

	wg.Wait()
	log.Println("Overall time: ", time.Since(start))
	log.Printf("Errors: %d\n", errors)
	log.Printf("Successes: %d\n", success)
	
}