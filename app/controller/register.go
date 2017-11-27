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
	// db, err := gorm.Open("mysql", "root:root@/library?charset=utf8&parseTime=True&loc=Local")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body",
			http.StatusInternalServerError)
	}
	resp, err := service.CreateUser([]byte(body))
	if err != nil {
		fmt.Fprintln(w, err)
	}
	b, err := json.Marshal(resp)
	if err != nil {
		fmt.Fprintln(w, err)
	}
	fmt.Fprintln(w, "response is = ", b)
	//post request with json in body using struct
	// var s = new(&Person)
	// person := new(model.Person)
	// json.Unmarshal([]byte(body), person)
	// db.Create(person)
	// fmt.Fprintln(w, "name is = ", person.Name)

	//post request with json in body using map
	// var name map[string]interface{}
	// json.Unmarshal([]byte(body), &name)
	// fmt.Fprintln(w, "name is = ", name["name"])

	//get request with queryparameter in url
	// fmt.Fprintf(w, "name is = %s", r.URL.Query().Get("name"))
}
