package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	pb "perScore/perScoreProto/usersDetails"

	"google.golang.org/grpc"
)

// GetEntrie ...
func GetEntrie(body []byte) (*pb.GetEntriesResponse, error) {
	ctx := context.Background()
	entries := new(pb.GetEntriesRequest)
	json.Unmarshal([]byte(body), entries)
	entries.AuthToken = GetToken("praveenraturi3@yahoo.com")
	conn, err := grpc.Dial(os.Getenv("PER_SCORE_CALC_SERVICE_DIAL"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial gRPC connection: %v", err)
	}
	defer conn.Close()

	questionClientConnection := pb.NewUserClient(conn)

	response, err := questionClientConnection.GetEntries(ctx, entries)
	fmt.Println("response", response)
	return response, err
}
