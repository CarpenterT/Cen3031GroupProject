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

type Server struct {
	gorm.Model

	ID              int64
	ServerCreatorID string `gorm:"<-"`
	ServerName      string `gorm:"<-"`
	ServerSize      int64  `gorm:"<-"`
}

type ServerGroups struct {
	gorm.Model

	ID       int64
	GroupID  int64 `gorm:"<-"`
	ServerID int64 `gorm:"<-"`
}

type ServerUsers struct {
	gorm.Model

	ID       int64
	UserID   int64 `gorm:"<-"`
	ServerID int64 `gorm:"<-"`
}

type Groups struct {
	gorm.Model

	ID             int64
	GroupCreatorID string
	GroupName      string `gorm:"<-"`
	GroupSize      int64  `gorm:"<-"`
}

type GroupMembers struct {
	gorm.Model

	UserID  int64
	GroupID int64
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
	db.AutoMigrate(&Groups{})
	db.AutoMigrate(&GroupMembers{})
	db.AutoMigrate(&Server{})
	db.AutoMigrate(&ServerGroups{})
	db.AutoMigrate(&ServerUsers{})

	return status
}

func main() {
	initalSQLDataBase()
}
