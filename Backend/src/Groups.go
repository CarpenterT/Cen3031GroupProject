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

// DB *gorm.DB
//var err error

//const DNS = "database\\ClusterC.db"

type Groups struct {
	ID             uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	GroupCreatorID string         `json:"Creator"`
	GroupName      string         `json:"groupname"`
	GroupSize      int64          `json:"groupsize"`
}

type GroupMembers struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	UserID    int64          `json:"userID"`
	GroupID   int64          `json:"groupID"`
}

// function to handle initial migration and opening
func InitialGroupsMigration() {
	DB, err = gorm.Open(sqlite.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&Groups{})
	DB.AutoMigrate(&GroupMembers{})
}

// function handles asking for a list of all groups
func GetGroups(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var groups []Groups
	DB.Find(&groups)
	json.NewEncoder(w).Encode(groups)
}

// function handles getting one particular group by ID
func GetGroupByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var groups Groups
	DB.First(&groups, params["id"])
	json.NewEncoder(w).Encode(groups)
}

// function handles getting one particular group by groupname
func GetGroupByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var groups Groups

	err := DB.Where(" groupname = ?", params["groupname"]).First(&groups).Error
	//If there was no error, we can send data
	if err == nil {
		json.NewEncoder(w).Encode("Group found.")
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		//if there was an error, like ErrRecordNotFound, we can return that it failed
		json.NewEncoder(w).Encode("Group not found.")
	} else {
		json.NewEncoder(w).Encode("Unknown Error.")
	}

}

// function handles creating a group
func CreateGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var groups Groups
	json.NewDecoder(r.Body).Decode(&groups)
	DB.Create(&groups)
	json.NewEncoder(w).Encode("Group successfully created.")
}

// function handles updating a group already in the database
func UpdateGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var groups Groups
	DB.First(&groups, params["id"])
	json.NewDecoder(r.Body).Decode(&groups)
	DB.Save(&groups)
	json.NewEncoder(w).Encode(groups)
}

// function soft-deletes a group in the database
func DeleteGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var groups Groups
	DB.Delete(&groups, params["id"])
	json.NewEncoder(w).Encode("The group is deleted successfully.")
}

// // function add a user to group
func AddGroupMember(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var groupMembers GroupMembers
	json.NewDecoder(r.Body).Decode(&groupMembers)
	DB.Create(&groupMembers)
	json.NewEncoder(w).Encode("User successfully added.")

}

// function removes a user from group
func DeleteGroupMember(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var groupMembers GroupMembers
	DB.Where("userID = ?", params["userID"]).Delete(&groupMembers)
	json.NewEncoder(w).Encode("The user is removed successfully.")
}
