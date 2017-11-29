package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"perScore/app/service"
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
	b, err := json.Marshal(resp)
	if err != nil {
		fmt.Fprintln(w, err)
	}
	fmt.Fprintln(w, string(b))
}
