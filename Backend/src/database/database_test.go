package main

import (
	"fmt"
	"testing"
)

func TestInitalSQLDataBase(t *testing.T) {

	got := initalSQLDataBase()
	want := "Success"

	if got != want {
		fmt.Println("Database Test Failed")
		t.Errorf("got %q, wanted %q", got, want)
	}

	if got == want {
		fmt.Println("Database Test Passed")
	}

}
