package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Server struct {
	ID         uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	ServerName string         `json:"name"`
	Admin      string         `json:"admin"`
}

type ServerGroups struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	GroupID   uint           `json:"groupID"`
	ServerID  uint           `json:"serverID"`
}

type ServerUsers struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	UserID    uint           `json:"userID"`
	ServerID  uint           `json:"serverID"`
}

// function to handle initial migration and opening
func InitialServerMigration() {
	DB, err = gorm.Open(sqlite.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&Server{})
	DB.AutoMigrate(&ServerGroups{})
	DB.AutoMigrate(&ServerUsers{})
}

// function handles asking for a list of all servers
func GetServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var server []Server
	DB.Find(&server)
	json.NewEncoder(w).Encode(server)
}

// function handles getting one particular server by ID
func GetServerByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var server Server
	DB.First(&server, params["id"])
	json.NewEncoder(w).Encode(server)
}

// function handles getting one particular server by name
func GetServerByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var server Server

	err := DB.Where("servername = ?", params["servername"]).First(&server).Error
	//If there was no error, we can send data
	if err == nil {
		json.NewEncoder(w).Encode("Server found.")
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		//if there was an error, like ErrRecordNotFound, we can return that it failed
		json.NewEncoder(w).Encode("Server not found.")
	} else {
		json.NewEncoder(w).Encode("Unknown Error.")
	}

}

// function handles creating a server
func CreateServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var server Server
	json.NewDecoder(r.Body).Decode(&server)
	DB.Create(&server)
	json.NewEncoder(w).Encode("Server successfully created.")
}

// function handles updating a server already in the database
func UpdateServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var server Server
	DB.First(&server, params["id"])
	json.NewDecoder(r.Body).Decode(&server)
	DB.Save(&server)
	json.NewEncoder(w).Encode(server)
}

// function soft-deletes a server in the database
func DeleteServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var server Server
	DB.Delete(&server, params["id"])
	json.NewEncoder(w).Encode("The server is deleted successfully.")
}

// function adds a group to server
func AddGroupToServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var servergroups ServerGroups
	json.NewDecoder(r.Body).Decode(&servergroups)
	DB.Create(&servergroups)
	json.NewEncoder(w).Encode("Group successfully added.")

}

// function removes a group from server
func DeleteGroupFromServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var servergroups ServerGroups
	DB.Where("groupID = ?", params["groupID"]).Delete(&servergroups)
	json.NewEncoder(w).Encode("The group is removed successfully.")
}

// function adds a user to server
func AddUserToServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var serverusers ServerUsers
	json.NewDecoder(r.Body).Decode(&serverusers)
	DB.Create(&serverusers)
	json.NewEncoder(w).Encode("User successfully added.")

}

// function removes a user from server
func DeleteUserFromServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var serverusers ServerUsers
	DB.Where("userID = ?", params["userID"]).Delete(&serverusers)
	json.NewEncoder(w).Encode("The user is removed successfully.")
}
