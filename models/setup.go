package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/crud_learning_vue?parseTime=true"))
	if err != nil {
		panic(err)
	}

	// database.AutoMigrate(&Student{})

	DB = database
}
