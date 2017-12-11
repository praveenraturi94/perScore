package service

import (
	"context"
	"fmt"
	"log"
	"os"
	pb "perScoreServer/perScoreProto/perScoreCal/user"

	"google.golang.org/grpc"
)

// GetEntries ...
func GetEntries(token string) (*pb.GetEntriesResponse, error) {
	ctx := context.Background()
	entriesRequest := new(pb.GetEntriesRequest)
	entriesRequest.AuthToken = token
	fmt.Println("entriesRequest", entriesRequest)

	// json.Unmarshal([]byte(body), interest)
	// entries.AuthToken = GetToken("praveenraturi3@yahoo.com")
	conn, err := grpc.Dial(os.Getenv("PER_SCORE_CALC_SERVICE_DIAL"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial gRPC connection: %v", err)
	}
	defer conn.Close()

	questionClientConnection := pb.NewUserClient(conn)

	response, err := questionClientConnection.GetEntries(ctx, entriesRequest)
	fmt.Println("GetEntries Response:", response)
	return response, err
}
