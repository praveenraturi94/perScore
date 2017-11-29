package service

import (
	"context"
	"encoding/json"
	"log"
	pb "perScore/perScoreProto/question"

	"google.golang.org/grpc"
)

// ListCategory ...
// type listCategory struct {
// 	Categories []category `json:"list_categories"`
// }
//
// type category struct {
// 	ID          int32  `json:"id"`
// 	Name        string `json:"name"`
// 	WeightRange int32  `json:"weight_range"`
// 	Parent      int32  `json:"parent"`
// 	Level       int32  `json:"level"`
// }

// CreateQuestion ...
func CreateQuestion(body []byte) (*pb.CreateQuestionResponse, error) {

	ctx := context.Background()

	ques := new(pb.CreateQuestionRequest)
	// categoryList := new(listCategory)
	json.Unmarshal([]byte(body), ques)

	// json.Unmarshal([]byte(body), categoryList)
	// fmt.Fprintln(w, "name is = ", name["name"])
	// var newCategories []*pb.CreateQuestionRequest_Category
	// newCategories := make([]*pb.CreateQuestionRequest_Category, len(categoryList.Categories))

	// for j := 0; j <= len(categoryList.Categories)-1; j++ {
	// 	categ := iterateCategories(ques.Categories)
	//
	// 	newCategories = append(newCategories, categ...)
	//
	// 	// for i := 0; i <= len(ques.Categories)-1; i++ {
	// 	// 	categ := createCategoryFamilyLine(ques.Categories[i])
	// 	// 	// fmt.Println("category family line ", categ)
	// 	// 	newCategories = append(newCategories, categ...)
	// 	// 	// fmt.Println("new category array ", newCategories)
	// 	// }
	// }

	// fmt.Println("newCategories ", newCategories)
	// ques.Categories = newCategories
	ques.AuthToken = GetToken("praveenraturi3@yahoo.com")
	conn, err := grpc.Dial("localhost:6060", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial gRPC connection: %v", err)
	}
	defer conn.Close()

	questionClientConnection := pb.NewQuestionClient(conn)

	response, err := questionClientConnection.CreateQuestion(ctx, ques)

	return response, err
}
