package service

import (
	"context"
	"encoding/json"
	"log"
	"perScore/perScoreProto/question"

	"google.golang.org/grpc"
)

// GetQuestion ...
func GetQuestion(body []byte) (*question.GetQuestionResponse, error) {
	ctx := context.Background()
	conn, err := grpc.Dial("localhost:6060", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial gRPC connection: %v", err)
	}
	defer conn.Close()

	requestQues := new(question.GetQuestionRequest)
	json.Unmarshal([]byte(body), requestQues)
	requestQues.AuthToken = GetToken("praveenraturi3@yahoo.com")
	questionClientConnection := question.NewQuestionClient(conn)

	response, err := questionClientConnection.GetQuestion(ctx, requestQues)
	return response, err
}
