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

var DB *gorm.DB
var err error

const DNS = "database\\ClusterC.db"

type User struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Username  string         `json:"username"`
	Password  string         `json:"password"`
}

// function to handle initial migration and opening
func InitialMigration() {
	DB, err = gorm.Open(sqlite.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&User{})
}

// function handles asking for a list of all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []User
	DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}

// function handles getting one particular user by ID
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	DB.First(&user, params["id"])
	json.NewEncoder(w).Encode(user)
}

// function handles getting one particular user by Username
func GetUserByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User

	err := DB.Where("username = ?", params["username"]).First(&user).Error
	//If there was no error, we can send data
	if err == nil {
		json.NewEncoder(w).Encode("User found.")
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		//if there was an error, like ErrRecordNotFound, we can return that it failed
		json.NewEncoder(w).Encode("User not found.")
	} else {
		json.NewEncoder(w).Encode("Unknown Error.")
	}

}

// function handles checking if a user has the right password
func CheckPass(w http.ResponseWriter, r *http.Request) {
	println("Reached")
	w.Header().Set("Content-Type", "application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	// Check if username AND password are both valid for a row.
	err := DB.Where("password = ? AND username = ?", user.Password, user.Username).First(&user).Error
	if err == nil {
		json.NewEncoder(w).Encode("Password validated.")
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		//if there was an error, like ErrRecordNotFound, we can return that it failed
		json.NewEncoder(w).Encode("Invalid.")
	} else {
		json.NewEncoder(w).Encode("Unknown Error.")
	}

}

// function handles creating a user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	DB.Create(&user)
	json.NewEncoder(w).Encode("User successfully created.")
}

// function handles updating a user already in the database
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	DB.First(&user, params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	DB.Save(&user)
	json.NewEncoder(w).Encode(user)
}

// function soft-deletes a user in the database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	DB.Delete(&user, params["id"])
	json.NewEncoder(w).Encode("The user is deleted successfully.")
}
