package main

import (
	"fmt"
	"net/http"
)

// portNumber is a constant that holds the port number for the application
const portNumber = ":8080" // http://localhost:8080/

//* Handler Functions
// Home is the home page handler function
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

// About is the about page handler function
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	fmt.Fprintf(w, "This is the about page and 2 + 2 is %d", sum)
}

//* Normal Functions
// addValues is a function that adds two given integers
func addValues(x int, y int) int {
	return x + y
}

// main is the main application function
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Starting server on port %s\n", portNumber)
	http.ListenAndServe(portNumber, nil)
}
