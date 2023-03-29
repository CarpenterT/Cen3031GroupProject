//go:build exclude

package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"<-"`
	Password string `gorm:"<-"`
}

func testGorm() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Print("Failed")
		panic("failed to connect database")
	}

	var user User

	db.AutoMigrate(&User{})

	db.Create(&User{Username: "username", Password: "sucess"})
	//db.Create(&User{Username: "username1", Password: "12346"})
	//db.Create(&User{Username: "username2", Password: "12345"})

	user = User{}
	db.First(&user, 1)
	//db.Where("Username = ?", "username1").First(&user)
	fmt.Println(user.Password)
	//db.Delete(&user, 1)
}
