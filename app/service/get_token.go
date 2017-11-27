package service

import (
		"github.com/jinzhu/gorm"
		"log"
		"perScore/app/model"
)
// GetToken ...
func GetToken(email string) string {

	db, err := gorm.Open("mysql", "root:root@/library?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	token := new(model.Token)
	err1 := db.Where("email = ?", email).First(&token)
	if err1.RecordNotFound() == true {

	} else {
		return token.Token
	}
	defer db.Close()

	return "null"
}
