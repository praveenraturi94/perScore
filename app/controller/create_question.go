package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"perScoreServer/app/service"
)

// CreateQuestion ...
func CreateQuestion(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body",
			http.StatusInternalServerError)
	}
	// Calling perScoreCal RPC
	response, _ := service.CreateQuestion(body)

	fmt.Println("CreateQuestion Response:", response)
	if err = json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}
