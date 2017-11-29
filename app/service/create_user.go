package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"perScore/perScoreProto/user"

	"google.golang.org/grpc"
)

// CreateUser ...
func CreateUser(body []byte) (*user.CreateUserResponse, error) {

	ctx := context.Background()

	conn, err := grpc.Dial("192.168.100.88:6060", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial gRPC connection: %v", err)
	}

	defer conn.Close()

	newUser := new(user.CreateUserRequest)
	json.Unmarshal([]byte(body), newUser)

	fmt.Println("user data contains = ", newUser)
	createUserClientConnection := user.NewUserClient(conn)
	response, err := createUserClientConnection.CreateUser(ctx, newUser)

	return response, err
}
