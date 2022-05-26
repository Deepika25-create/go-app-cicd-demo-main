package main

import (
	"fmt"
	"net/http"
)

func main() {
	// create a route handler to listen on "/"
	http.HandleFunc("/", Home)
	fmt.Println("Starting server on port :8080")
	// Start the server
	http.ListenAndServe(":8080", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	// Assign a message to return
	msg := "This is a test for a toolchain"

	// Write this message to the response to display on the screen
	w.Write([]byte(fmt.Sprintf(msg)))
}
