package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	pb "perScoreServer/perScoreProto/user"

	"google.golang.org/grpc"
)

// CreateUser ...
func CreateUser(request []byte) (*pb.CreateUserResponse, error) {
	ctx := context.Background()
	conn, err := grpc.Dial(os.Getenv("PER_SCORE_AUTH_SERVICE_DIAL"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial gRPC connection: %v", err)
	}
	defer conn.Close()

	newUser := new(pb.CreateUserRequest)
	json.Unmarshal([]byte(request), newUser)
	fmt.Println("Request = ", string(request))
	fmt.Println("User = ", *newUser)
	fmt.Println("Location = ", *newUser.Location)
	clientConnection := pb.NewUserClient(conn)
	response, err := clientConnection.CreateUser(ctx, newUser)
	fmt.Println("Response = ", response)
	return response, err
}
