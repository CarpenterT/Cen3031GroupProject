package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// database table sturctures

type User struct {
	gorm.Model
	Username string `gorm:"<-"`
	Password string `gorm:"<-"`
}

func initalSQLDataBase() string {
	db, err := gorm.Open(sqlite.Open("ClusterC.db"), &gorm.Config{})
	var status string
	if err != nil {
		status = "Failed"
		panic("failed to connect database")
	}

	status = "Success"
	db.AutoMigrate(&User{})

	return status
}

func main() {
	initalSQLDataBase()
}
