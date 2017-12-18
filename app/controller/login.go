package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"perScoreServer/app/service"
	pbc "perScoreServer/perScoreProto/perScoreCal/user"
	pba "perScoreServer/perScoreProto/user"

	log "github.com/sirupsen/logrus"
)

// LoginResponse ...
type LoginResponse struct {
	Response *pba.GetSessionResponse
	Token    string
}

// GetEntriesResponse ...
type GetEntriesResponse struct {
	Response *pbc.GetEntriesResponse
	Token    string
}

//Login ...
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request:", r)
	params, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Error reading request body",
			http.StatusInternalServerError)
	} else {
		fmt.Println("Params:", params)
	}

	response, err := service.Login(params)
	if err != nil {
		log.Errorf("Error in login: %v", err)
	}
	token := response.Token

	if response.Status == "FAILURE" {
		loginResponse := new(LoginResponse)
		loginResponse.Response = response
		if err = json.NewEncoder(w).Encode(loginResponse); err != nil {
			panic(err)
		}
	} else {
		responseGetEntries, _ := service.GetEntries(token)
		getEntriesResponse := new(GetEntriesResponse)
		getEntriesResponse.Response = responseGetEntries
		getEntriesResponse.Token = token
		fmt.Println("GetEntries Response:", getEntriesResponse)
		if err = json.NewEncoder(w).Encode(getEntriesResponse); err != nil {
			panic(err)
		}
	}
}
