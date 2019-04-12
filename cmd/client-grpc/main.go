package main

import (
	"context"
	"flag"
	"log"
	v1 "protobuf-grpc-http-tutorial/go-grpc-http-rest-microservice-tutorial/pkg/api/v1"
	"time"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
)

const (
	apiVersion = "v1"
)

func main() {
	address := flag.String("server", "", "gRPC server in format host:port")
	flag.Parse()
	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := v1.NewToDoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	t := time.Now().In(time.UTC)
	reminder, _ := ptypes.TimestampProto(t)
	pfx := t.Format(time.RFC3339Nano)

	// Call Create
	req1 := v1.CreateRequest{
		Api: apiVersion,
		ToDo: &v1.ToDo{
			Title:       "title (" + pfx + ")",
			Description: "description (" + pfx + ")",
			Reminder:    reminder,
		},
	}
	res1, err := c.Create(ctx, &req1)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create result: <%+v>\n\n", res1)

	id := res1.Id
	log.Printf("id is %v", id)
}
