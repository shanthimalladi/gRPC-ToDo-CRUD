package grpc

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	v1 "protobuf-grpc-http-tutorial/go-grpc-http-rest-microservice-tutorial/pkg/api/v1"

	"google.golang.org/grpc"
)

//RunServer runs gRPC service to publish TODO service
func RunServer(ctx context.Context, v1API v1.ToDoServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	v1.RegisterToDoServiceServer(server, v1API)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	// starting a go routine to exit server gracefully in case of os interrupts
	go func() {
		for range c {
			log.Println("shutting down gRPC server")
			server.GracefulStop()
			<-ctx.Done()
		}
	}()

	log.Println("Starting gRPC server")
	return server.Serve(listen)
}


