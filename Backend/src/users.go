package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type user struct {
	username string `json:"username"`

	password string `json:"password"`
}

func handleUserGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	page := vars["page"]

	fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
}

func handleUserPost(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	username := vars["username"]
	password := vars["password"]

	addUser(username, password)

	fmt.Fprintf(w, "You've requested the book: %s on page %s\n", username, password)

}

func addUser(username string, password string) {

	db, err := gorm.Open(sqlite.Open("ClusterC.db"), &gorm.Config{})
	if err != nil {
		fmt.Print("Failed")
		panic("failed to connect database")
	}

	db.Create(&User{Username: username, Password: password})
}
