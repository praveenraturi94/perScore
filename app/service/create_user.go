package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"perScore/perScoreProto/user"

	"google.golang.org/grpc"
)

// CreateUser ...
func CreateUser(body []byte) (*user.CreateUserResponse, error) {

	ctx := context.Background()

	conn, err := grpc.Dial(os.Getenv("PER_SCORE_AUTH_SERVICE_DIAL"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial gRPC connection: %v", err)
	}

	defer conn.Close()

	newUser := new(user.CreateUserRequest)
	json.Unmarshal([]byte(body), newUser)
	fmt.Println("res data = ", string(body))
	fmt.Println("user data contains = ", newUser)
	createUserClientConnection := user.NewUserClient(conn)
	response, err := createUserClientConnection.CreateUser(ctx, newUser)
	fmt.Println("response", response)
	return response, err
}
