package service

import (
	"context"
	"encoding/json"
	"log"
	"os"
	pb "perScore/perScoreProto/question"

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
func CreateQuestion(body []byte) (*pb.CreateQuestionRequest_Category, error) {
	ctx := context.Background()
	ques := new(pb.CreateQuestionRequest)
	json.Unmarshal([]byte(body), ques)
	ques.AuthToken = GetToken("praveenraturi3@yahoo.com")
	conn, err := grpc.Dial(os.Getenv("PER_SCORE_CALC_SERVICE_DIAL"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial gRPC connection: %v", err)
	}
	defer conn.Close()

	questionClientConnection := pb.NewQuestionClient(conn)

	response, err := questionClientConnection.CreateQuestion(ctx, ques)
	return Sorting(response), err
}

var categ = new(pb.CreateQuestionRequest_Category)
var cateTemp = new(pb.CreateQuestionResponse_Category)

// Sorting ...
func Sorting(response *pb.CreateQuestionResponse) *pb.CreateQuestionRequest_Category {
	categories := response.Categories
	for _, category := range categories {
		level := category.GetLevel()
		if level == 1 {
			var categorey = new(pb.CreateQuestionRequest_Category)
			categorey.Id = category.GetId()
			categorey.Name = category.GetName()
			categorey.Parent = category.GetParent()
			categ.Categories = append(categ.Categories, categorey)
		}
	}
	for {
		for i, category := range categories {
			for _, cate := range categ.Categories {
				if category.GetParent() == cate.GetId() {
					var categorey = new(pb.CreateQuestionRequest_Category)
					categorey.Id = category.GetId()
					categorey.Name = category.GetName()
					categorey.Parent = category.GetParent()
					cate.Categories = append(cate.Categories, categorey)

				}
			}
			if i >= len(categories) {
				break
			} else {
				categories = append(categories[:i], categories[i+1:]...)
			}
		}
		if len(categories) == 0 {
			break
		}
	}
	return categ
}
