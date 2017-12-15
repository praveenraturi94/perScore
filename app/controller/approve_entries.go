package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"perScoreServer/app/service"
)

// ApproveEntries ...
func ApproveEntries(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body",
			http.StatusInternalServerError)
	}
	response, err := service.ApproveEntrie(body)
	if err != nil {
		fmt.Fprintln(w, "Error while creating question ", err)
	}
	if err = json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Fprintln(w, string(b))
}
