package model

import (
	// _ "github.com/go-sql-driver/mysql"

	"fmt"
	"perScore/app/shared"

	_ "github.com/jinzhu/gorm/dialects/mysql" // ...
)

//SetupDatabase ...
func SetupDatabase() error {
	fmt.Println("migration")

	err := shared.DBCon.AutoMigrate(
		&Token{},
	).Error
	return err
}
