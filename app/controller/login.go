package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"perScore/app/service"
)

//Login ...
func Login(w http.ResponseWriter, r *http.Request) {
	// db, err := gorm.Open("mysql", "root:root@/library?charset=utf8&parseTime=True&loc=Local")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body",
			http.StatusInternalServerError)
	}

	// b, err := json.Marshal(service.Login(body))
	// if err != nil {
	// 	fmt.Fprintln(w, err)
	// }
	token, err := service.Login(body)
	if err != nil {
		fmt.Fprintln(w, "Error while logging in = ", err)
	}
	fmt.Fprintln(w, "response is = ", token)
	//post request with json in body using struct
	// var s = new(&Person)
	// person := new(model.Person)
	// json.Unmarshal([]byte(body), person)
	// err1 := db.Where("email = ?", person.Email).First(&person)
	// if err1.RecordNotFound() == true {
	// 	fmt.Fprintln(w, "record not found")
	// } else {
	// 	fmt.Fprintln(w, "record found")
	// }
}
