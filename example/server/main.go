package main

import (
	"context"
	"log"
	"net"

	user_pb "github.com/shubham-rewale/grpc-unary-example/gen/go/user/v1"

	"google.golang.org/grpc"
)

type userService struct{}

func (u *userService) GetUser(_ context.Context, req *user_pb.GetUserRequest) (*user_pb.GetUserResponse, error) {
	return &user_pb.GetUserResponse{
		User: &user_pb.User{
			Uuid:     req.Uuid,
			FullName: "Mario",
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalf("Failed to Listen with error %v", err)
	}

	grpcServer := grpc.NewServer()

	user_pb.RegisterUserServiceServer(grpcServer, &userService{})

	grpcServer.Serve(lis)
}
