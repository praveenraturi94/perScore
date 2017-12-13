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
	response, err := service.CreateQuestion(body)
	if err != nil {
		fmt.Fprintln(w, "Error while creating question ", err)
	}
	if err = json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}
