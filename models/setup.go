package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:1234567890@tcp(localhost:3306)/basic_crud"
	// database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/basic_crud"))
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Product{})

	DB = database
}
