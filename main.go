package main

import (
	"fmt"
	"net/http"
)

// portNumber is a constant that holds the port number for the application
const portNumber = ":8080" // http://localhost:8080/

//* Main Function
// main is the main application function
func main() {
	// call the http.HandleFunc function to register the Home and About handlers
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	// call the http.ListenAndServe function to start the server
	fmt.Printf("Starting server on port %v\n", portNumber)
	http.ListenAndServe(portNumber, nil)
}
