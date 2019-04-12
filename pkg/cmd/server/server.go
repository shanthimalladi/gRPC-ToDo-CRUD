package cmd

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"protobuf-grpc-http-tutorial/go-grpc-http-rest-microservice-tutorial/pkg/protocol/grpc"
	v1 "protobuf-grpc-http-tutorial/go-grpc-http-rest-microservice-tutorial/pkg/service/v1"
)

//Config is the configuration for server
type Config struct {
	GRPCPort   string
	DBHost     string
	DBUser     string
	DBPassword string
	DBSchema   string
}

//RunServer runs gRPC server and HTTP gateway
func RunServer() error {
	ctx := context.Background()

	var cfg Config
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "", "grpc port to bind")
	flag.StringVar(&cfg.DBHost, "grpc-host", "", "grpc host")
	flag.StringVar(&cfg.DBUser, "grpc-user", "", "grpc user")
	flag.StringVar(&cfg.DBPassword, "grpc-password", "", "grpc password")
	flag.StringVar(&cfg.DBSchema, "grpc-schema", "", "grpc schema")
	param := "parseTime=true"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)%s?%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBSchema, param)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	v1API := v1.NewToDoServiceServer(db)
	return grpc.RunServer(ctx, v1API, cfg.GRPCPort)
}
