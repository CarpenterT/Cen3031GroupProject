//go:build exclude

package main

import (
	"log"
	"net/http"
)

func main() {

	testGorm()

	host := "127.0.0.1:8080"
	if err := http.ListenAndServe(host, httpHandler()); err != nil {
		log.Fatalf("Failed to listen on %s: %v", host, err)
	}

}
