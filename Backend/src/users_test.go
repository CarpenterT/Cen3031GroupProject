package main

import (
	"fmt"
	"testing"
)

func TestAddUser(t *testing.T) {

	var user string
	var pass string
	user = "Ross"
	pass = "Success"
	got := addUser(user, pass)
	want := "Success"

	if got != want {
		fmt.Println("Database Test Failed")
		t.Errorf("got %q, wanted %q", got, want)
	}

	if got == want {
		fmt.Println("Database Test Passed")
	}

}
