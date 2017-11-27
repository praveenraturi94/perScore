package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"perScore/app/service"
)

// GetQues ...
func GetQues(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body",
			http.StatusInternalServerError)
	}

	b, err := json.Marshal(service.GetQuestion(body))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintln(w, string(b))
	// fmt.Fprintln(w, "response is = ", service.GetQuestion(body))

}
