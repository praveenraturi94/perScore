package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // ...
)

// Person ...
type Person struct {
	gorm.Model
	Name     string     `json:"name"`
	Age      int16      `json:"age"`
	Email    string     `json:"email"`
	Password string     `json:"password"`
	Location []Location `json:"location" gorm:"ForeignKey:PersonID"`
}

// Location ...
type Location struct {
	gorm.Model
	City     string `json:"city"`
	Country  string `json:"country"`
	PersonID uint
}
