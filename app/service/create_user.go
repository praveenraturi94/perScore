package service

import (
	"context"
	"encoding/json"
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
	clientConnection := pb.NewUserClient(conn)
	response, err := clientConnection.CreateUser(ctx, newUser)
	return response, err
}
