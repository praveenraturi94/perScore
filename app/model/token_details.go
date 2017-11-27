package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // ...
)

// Token ...
type Token struct {
	gorm.Model
	Token          string `json:"name"`
	ExpirationTime int    `json:"age"`
	Email          string
}
