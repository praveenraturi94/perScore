package service

import (
	"context"
	"encoding/json"
	"log"
	"os"
	pb "perScore/perScoreProto/usersDetails"

	"google.golang.org/grpc"
)

// GetInterest ...
func GetInterest(body []byte) (*pb.GetInterestResponse, error) {
	ctx := context.Background()
	interest := new(pb.GetInterestRequest)
	json.Unmarshal([]byte(body), interest)
	interest.AuthToken = GetToken("praveenraturi3@yahoo.com")
	conn, err := grpc.Dial(os.Getenv("PER_SCORE_CALC_SERVICE_DIAL"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial gRPC connection: %v", err)
	}
	defer conn.Close()

	questionClientConnection := pb.NewUserClient(conn)

	response, err := questionClientConnection.GetInterests(ctx, interest)

	return response, err
}
