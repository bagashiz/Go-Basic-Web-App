package main

import (
	"errors"
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
	// call the addValues function
	sum := addValues(2, 2)
	fmt.Fprintf(w, "This is the about page and 2 + 2 is %d", sum)
}

func Divide(w http.ResponseWriter, r *http.Request) {
	// call the divideValues function
	f, err := divideValues(100, 0)
	// check for error
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
		return // exit function
	}

	fmt.Fprintf(w, "The result is %f", f)
}

//* Normal Functions
// addValues is a function that adds two given integers
func addValues(x, y int) int {
	return x + y
}

// divideValues is a function that divides two given floats
func divideValues(x, y float64) (float64, error) {
	// Check for divide by zero
	if y == 0 {
		err := errors.New("CANNOT DIVIDE BY ZERO")
		return 0, err
	}
	result := x / y
	return result, nil
}

//* Main Function
// main is the main application function
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Printf("Starting server on port %s\n", portNumber)
	http.ListenAndServe(portNumber, nil)
}
