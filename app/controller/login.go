package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"perScoreServer/app/service"

	"github.com/gorilla/sessions"
	log "github.com/sirupsen/logrus"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

type LoginResponse struct {
	Response *pb.GetEntriesResponse
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

	mappedResult, loginResponse, err := service.Login(params)

	session, _ := store.Get(r, "session")
	// Set some session values.
	session.Values["email"] = mappedResult["email"]
	// Save it before we write to the response/return from the handler.
	session.Save(r, w)

	received := false
	if err != nil {
		log.Errorf("Error in login: %v", err)
	} else {
		fmt.Println("Token", loginResponse.Token)
		response, _ := service.GetEntries(loginResponse.Token)
		responseLogin := new(LoginResponse)
		responseLogin.Response = response
		responseLogin.Token = loginResponse.Token
		fmt.Println("response", response)
		if err = json.NewEncoder(w).Encode(responseLogin); err != nil {
			panic(err)
		}
		received = true
	}

	if received == false {
		if err = json.NewEncoder(w).Encode(loginResponse); err != nil {
			panic(err)
		}
	}
}
