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

var chatDB *gorm.DB

var chatErr error

const chatDNS = "database\\chat.db"

// Messages should contain an ID, when they were sent, who sent them,
//
//	and what the message says.
type Message struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Username  string         `json:"username"`
	Msg       string         `json:"msg"`
}

// function to handle initial migration and opening
func InitChatDB() {
	chatDB, err = gorm.Open(sqlite.Open(chatDNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	chatDB.AutoMigrate(&Message{})
	fmt.Println("Chat.db opened.")
}

// function handles asking for a list of all messages
func GetAllMsg(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var msgs []Message
	chatDB.Find(&msgs)
	json.NewEncoder(w).Encode(msgs)
}

// function handles getting one particular message by ID
func GetMsgByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var msg Message
	chatDB.First(&msg, params["id"])
	json.NewEncoder(w).Encode(msg)
}

// function handles finding a message's id by checking the time, user, and message text.
func GetMsgID(w http.ResponseWriter, r *http.Request) {
	//println("Reached")
	w.Header().Set("Content-Type", "application/json")
	var msg Message
	json.NewDecoder(r.Body).Decode(&msg)
	// Check if time, username, and message are valid for a row.
	err := chatDB.Where("created_at = ? AND username = ? AND msg = ?", msg.CreatedAt, msg.Username, msg.Msg).First(&msg).Error
	if err == nil {
		// If the message was found, return it's ID
		json.NewEncoder(w).Encode(msg.ID)
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		//if there was an error, like ErrRecordNotFound, we can return that it failed
		json.NewEncoder(w).Encode("Message not found.")
	} else {
		json.NewEncoder(w).Encode("Unknown Error.")
	}

}

// function handles creating a message
func CreateMsg(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var msg Message
	json.NewDecoder(r.Body).Decode(&msg)
	chatDB.Create(&msg)
	json.NewEncoder(w).Encode("Message sent.")
}

// function handles updating a message already in the database
func UpdateMsg(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var msg Message
	DB.First(&msg, params["id"])
	json.NewDecoder(r.Body).Decode(&msg)
	chatDB.Save(&msg)
	json.NewEncoder(w).Encode(msg)
}

// function soft-deletes a message in the database
func DeleteMsg(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var msg Message
	chatDB.Delete(&msg, params["id"])
	json.NewEncoder(w).Encode("The message was deleted.")
}
