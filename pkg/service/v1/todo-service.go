package v1

import (
	"context"
	"database/sql"
	"log"
	v1 "protobuf-grpc-http-tutorial/go-grpc-http-rest-microservice-tutorial/pkg/api/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	apiVersion = "v1"
)

type toDoServiceServer struct {
	db *sql.DB
}

//NewToDoServiceServer : Initializes a server
func NewToDoServiceServer(db *sql.DB) v1.ToDoServiceServer {
	return &toDoServiceServer{db: db}
}

func (s *toDoServiceServer) checkAPI(api string) error {
	if apiVersion != api {
		return status.Errorf(codes.Unimplemented, "Unsupported API version")
	}
	return nil

}

func (s *toDoServiceServer) connect(ctx context.Context) (*sql.Conn, error) {
	log.Printf("Received connect")
	c, err := s.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database ->"+err.Error())
	}
	return c, nil
}

func (s *toDoServiceServer) Create(ctx context.Context, req *v1.CreateRequest) (*v1.CreateResponse, error) {
	if err:=s.checkAPI(req.Api); err!=nil{
		return nil, err
	}

	return &v1.CreateResponse{
		Api:apiVersion,
		Id: 20
	}, nil


}

func (s *toDoServiceServer) Update(ctx context.Context, req *v1.UpdateRequest) (*v1.UpdateResponse, error) {
	if err:=s.checkAPI(req.Api); err!=nil{
		return nil, err
	}

	return &v1.UpdateResponse{
		Api:apiVersion,
		Updated:3,
	}, nil
}

func (s *toDoServiceServer) Read(ctx context.Context, req *v1.ReadRequest) (*v1.ReadResponse, error) {
	if err:=s.checkAPI(req.Api); err!=nil{
		return nil, err
	}
	return &v1.ReadResponse{
		Api: apiVersion,
		ToDo: v1.ToDo{
			Id:1,
			Title: "shanthi",
			Description:"Dummy shanthi description",
			Reminder: time.Now().In(time.UTC),
		}
	}, nil
}

func (s *toDoServiceServer) Delete(ctx context.Context, req *v1.DeleteRequest) (*v1.DeleteResponse, error) {
	if err:=s.checkAPI(req.Api); err!=nil{
		return nil, err
	}
	return &DeleteResponse{
		Api:apiVersion,
		Deleted:20

	}
}
