package service

import (
	"context"
	"encoding/json"
	"log"
	"perScore/perScoreProto/question"

	"google.golang.org/grpc"
)

// GetQuestion ...
func GetQuestion(body []byte) *question.GetQuestionResponse {
	ctx := context.Background()
	conn, err := grpc.Dial("Vikram-Anand.local:6060", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial gRPC connection: %v", err)
	}
	defer conn.Close()

	requestQues := new(question.GetQuestionRequest)
	json.Unmarshal([]byte(body), requestQues)
	requestQues.AuthToken = GetToken("praveenraturi3@yahoo.com")
	questionClientConnection := question.NewQuestionClient(conn)

	response, err := questionClientConnection.GetQuestion(ctx, requestQues)
	if err != nil {
		log.Fatal("unable to fetch questions", err)
	}
	return response
}
