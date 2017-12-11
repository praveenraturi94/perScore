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
	session, _ := store.Get(r, "session")
	email := session.Values["email"].(string)
	response, err := service.CreateQuestion(body, email)
	if err != nil {
		fmt.Fprintln(w, "Error while creating question ", err)
	}
	b, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintln(w, string(b))
}
