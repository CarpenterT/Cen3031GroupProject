package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func initializeRouter() {

	r := mux.NewRouter()

	//defining the functions needed to handle users
	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", GetUserByID).Methods("GET")
	r.HandleFunc("/users/user/{username}", GetUserByName).Methods("GET")
	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/users/user", CheckPass).Methods("POST")
	r.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

	//this line fixes issues with cross origin garbage. Finally.
	handler := cors.AllowAll().Handler(r)

	log.Fatal(http.ListenAndServe(":8080", handler))
}

// To run, cd into the backend/src folder, type "go build", then ./m.exe or whatever the executable is named for you.
// This gets the server running, and in the web app you can now add to the database.
// We are writing to ClusterC.db in backend/src/database. Do not hard-code your full path. See user.go.
func main() {
	InitialMigration()
	initializeRouter()

	//unit tests go here
}
