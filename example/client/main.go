package main

import (
	"context"
	"fmt"
	"log"

	user_pb "github.com/shubham-rewale/grpc-unary-example/gen/go/user/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	conn, err := grpc.Dial("127.0.0.1:8080", opts...)

	if err != nil {
		log.Fatalf("Failed to dial with error %v", err)
	}

	defer conn.Close()

	client := user_pb.NewUserServiceClient(conn)

	res, err := client.GetUser(context.Background(), &user_pb.GetUserRequest{Uuid: "420"})

	if err != nil {
		log.Fatalf("fail to GetUser: %v", err)
	}

	fmt.Printf("This is the response %v", res)
}
