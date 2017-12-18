package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"perScoreServer/app/service"
)

// GetQuestion ...
func GetQuestion(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request", r)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body",
			http.StatusInternalServerError)
	}
	response, err := service.GetQuestion(body)
	if err != nil {
		fmt.Fprintln(w, "Error while fetching questions", err)
	}
	b, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintln(w, string(b))
}
