package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// portNumber is a constant that holds the port number for the application
const portNumber = ":8080" // http://localhost:8080/

//* Handler Functions
// Home is the home page handler function
func Home(w http.ResponseWriter, r *http.Request) {
	// call the renderTemplate function to render the home page
	renderTemplate(w, "home.html")
}

// About is the about page handler function
func About(w http.ResponseWriter, r *http.Request) {
	// call the renderTemplate function to render the about page
	renderTemplate(w, "about.html")
}

// renderTemplate is a function that renders the specified template to the response writer
func renderTemplate(w http.ResponseWriter, html string) {
	// Load the template from the templates folder and parse it into a html template object
	parsedTemplate, _ := template.ParseFiles("./templates/" + html)
	err := parsedTemplate.Execute(w, nil)
	// check for any errors
	if err != nil {
		fmt.Printf("Error parsing template: %v\n", err)
	}
}

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
