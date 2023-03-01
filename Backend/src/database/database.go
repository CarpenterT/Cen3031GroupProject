package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// database table sturctures

type User struct {
	gorm.Model
	Username string `gorm:"<-"`
	Password string `gorm:"<-"`
}

func initalSQLDataBase() {
	db, err := gorm.Open(sqlite.Open("ClusterC.db"), &gorm.Config{})
	if err != nil {
		fmt.Print("Failed")
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})
}

func main() {
	initalSQLDataBase()
}
