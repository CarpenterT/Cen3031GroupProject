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
	ServerName string `gorm:"<-"`
	Admin      string `gorm:"<-"`
}

type ServerGroups struct {
	gorm.Model
	GroupID  int64 `gorm:"<-"`
	ServerID int64 `gorm:"<-"`
}

type ServerUsers struct {
	gorm.Model
	UserID   int64 `gorm:"<-"`
	ServerID int64 `gorm:"<-"`
}

type Groups struct {
	gorm.Model
	GroupCreatorID string
	GroupName      string `gorm:"<-"`
	GroupSize      int64  `gorm:"<-"`
}

type Chat struct {
    gorm.Model
    Body string gorm:"<-"
    serverID int64 gorm:"<-"
    UserID int64 gorm:"<-"
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
  db.AutoMigrate(&Chat{})

	return status
}

func main() {
	print(initalSQLDataBase())
}
