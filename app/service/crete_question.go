package service

import (
	"context"
	"encoding/json"
	"fmt"
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
	conn, err := grpc.Dial("Vikram-Anand.local:6060", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial gRPC connection: %v", err)
	}
	defer conn.Close()

	questionClientConnection := pb.NewQuestionClient(conn)

	response, err := questionClientConnection.CreateQuestion(ctx, ques)
	Sorting(response)
	return response, err
}

var categ = new(pb.CreateQuestionRequest_Category)
var cateTemp = new(pb.CreateQuestionResponse_Category)

// Sorting ...
func Sorting(response *pb.CreateQuestionResponse) {
	fmt.Println("1")
	categories := response.Categories
	for _, category := range categories {
		level := category.GetLevel()
		// parent := category.Parent
		if level == 1 {
			var categorey = new(pb.CreateQuestionRequest_Category)
			categorey.Id = category.GetId()
			categorey.Name = category.GetName()
			categorey.Parent = category.GetParent()
			categ.Categories = append(categ.Categories, categorey)
		}
	}
	fmt.Println("2")
	fmt.Println("categ", categ)
	for {
		for _, category := range categories {
			for _, cate := range categ.Categories {
				if category.GetParent() == cate.GetId() {
					var categorey *pb.CreateQuestionRequest_Category
					categorey.Id = category.GetId()
					categorey.Name = category.GetName()
					categorey.Parent = category.GetParent()
					cate.Categories = append(cate.Categories, categorey)
					// categories.Categories = append(categories.Categories[:i], categories.Categories[i+1:]...)
					// cateTemp = append(cateTemp, categories)
				}
			}
		}
		if len(categories.Categories) == 0 {
			break
		}
	}
	fmt.Println(categ)
}
