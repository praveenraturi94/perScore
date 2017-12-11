package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres dialect for gorm
)

// Token ...
type Token struct {
	gorm.Model
	Token          string
	ExpirationTime int
	Email          string
}
