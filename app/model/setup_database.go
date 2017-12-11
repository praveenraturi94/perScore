package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

//SetupDatabase ...
func SetupDatabase(db *gorm.DB) error {
	fmt.Println("migration")

	err := db.AutoMigrate(
		&Token{},
	).Error
	return err
}
