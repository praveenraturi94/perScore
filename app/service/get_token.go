package service

import (
	"fmt"
	"log"
	"os"
	"perScoreServer/app/model"

	"github.com/jinzhu/gorm"
)

// GetToken ...
func GetToken(email string) string {

	dbString := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=%s", os.Getenv("DEV_HOST"), os.Getenv("DEV_DBNAME"), os.Getenv("DEV_USERNAME"), os.Getenv("DEV_PASSWORD"), os.Getenv("DEV_SSLMODE"))
	db, err := gorm.Open(os.Getenv("DEV_DB_DRIVER"), dbString)
	defer db.Close()
	if err != nil {
		log.Fatalf("Failed to create db connection: %v", err)
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
