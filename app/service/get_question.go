package service

import (
	"context"
	"encoding/json"
	"log"
	"os"
	pb "perScoreServer/perScoreProto/perScoreCal/question"

	"google.golang.org/grpc"
)

// GetQuestion ...
func GetQuestion(body []byte) (*pb.GetQuestionResponse, error) {
	ctx := context.Background()
	conn, err := grpc.Dial(os.Getenv("PER_SCORE_CALC_SERVICE_DIAL"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial gRPC connection: %v", err)
	}
	defer conn.Close()
	getQuestionRequest := new(pb.GetQuestionRequest)
	json.Unmarshal(body, getQuestionRequest)
	questionClientConnection := pb.NewQuestionClient(conn)

	response, err := questionClientConnection.GetQuestion(ctx, getQuestionRequest)
	return response, err
}
