package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"perScore/app/service"
)

// GetEntries ...
func GetEntries(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get entries")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body",
			http.StatusInternalServerError)
	}
	fmt.Println("get entries")
	response, err := service.GetEntrie(body)
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
