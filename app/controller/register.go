package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"perScoreServer/app/service"
)

//Register ...
func Register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("req", r)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body",
			http.StatusInternalServerError)
	}
	resp, err := service.CreateUser(body)
	if err != nil {
		fmt.Fprintln(w, "Error creating users = ", err)
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		panic(err)
	}
}
