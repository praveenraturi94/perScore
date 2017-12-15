package service

import (
	"context"
	"encoding/json"
	"log"
	"os"
	pb "perScoreServer/perScoreProto/perScoreCal/question"

	"google.golang.org/grpc"
)

// Category ...
type Category struct {
	Name       string
	Parent     uint
	Level      int32
	Approved   bool
	Categories []Category
}

// CreateQuestion ...
func CreateQuestion(body []byte) (*pb.CreateQuestionResponse, error) {
	ctx := context.Background()
	questionRequest := new(pb.CreateQuestionRequest)
	json.Unmarshal([]byte(body), questionRequest)
	conn, err := grpc.Dial(os.Getenv("PER_SCORE_CALC_SERVICE_DIAL"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial gRPC connection: %v", err)
	}
	defer conn.Close()

	questionClientConnection := pb.NewQuestionClient(conn)

	response, err := questionClientConnection.CreateQuestion(ctx, questionRequest)
	return response, err
}

var questionResponse = new(pb.CreateQuestionResponse)

// Sorting ...
func Sorting(response *pb.CreateQuestionResponse) *pb.CreateQuestionResponse {
	for _, category := range response.Categories {
		level := category.GetLevel()
		if level == 1 {
			var responseCategory = new(pb.CreateQuestionResponse_Category)
			responseCategory.Id = category.GetId()
			responseCategory.Name = category.GetName()
			responseCategory.Parent = category.GetParent()
			questionResponse.Categories = append(questionResponse.Categories, responseCategory)
		}
	}
	for {
		for i, category := range response.Categories {
			for _, questionResponseCategory := range questionResponse.Categories {
				if category.GetParent() == questionResponseCategory.GetId() {
					var responseCategory = new(pb.CreateQuestionResponse_Category)
					responseCategory.Id = category.GetId()
					responseCategory.Name = category.GetName()
					responseCategory.Parent = category.GetParent()
					questionResponse.Categories = append(questionResponse.Categories, responseCategory)
				}
			}
			if i >= len(response.Categories) {
				break
			} else {
				response.Categories = append(response.Categories[:i], response.Categories[i+1:]...)
			}
		}
		if len(response.Categories) == 0 {
			break
		}
	}
	return questionResponse
}
