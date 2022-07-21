package main

import (
	"net/http"
)

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
