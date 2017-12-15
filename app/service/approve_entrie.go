package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	pb "perScoreServer/perScoreProto/perScoreCal/user"

	"google.golang.org/grpc"
)

// ApproveEntrie ...
func ApproveEntrie(body []byte) (*pb.ApproveEntriesResponse, error) {
	ctx := context.Background()
	entries := new(pb.ApproveEntriesRequest)
	json.Unmarshal([]byte(body), entries)
	fmt.Println("entries", entries)
	conn, err := grpc.Dial(os.Getenv("PER_SCORE_CALC_SERVICE_DIAL"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial gRPC connection: %v", err)
	}
	defer conn.Close()

	questionClientConnection := pb.NewUserClient(conn)

	response, err := questionClientConnection.ApproveEntries(ctx, entries)

	return response, err
}
