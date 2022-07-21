package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bagashiz/Go-Basic-Web-App/pkg/config"
	"github.com/bagashiz/Go-Basic-Web-App/pkg/handlers"
	"github.com/bagashiz/Go-Basic-Web-App/pkg/render"
)

// portNumber is a constant that holds the port number for the application
const portNumber = ":8080" // http://localhost:8080/

//* Main Function
// main is the main application function
func main() {
	// create a new config object
	var app config.AppConfig
	// create a new template cache
	tc, err := render.CreateTemplateCache()
	// check for any errors
	if err != nil {
		log.Fatal("Error creating template cache: ", err)
	}
	// set the template cache
	app.TemplateCache = tc

	// call the http.HandleFunc function to register the Home and About handlers
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	// call the http.ListenAndServe function to start the server
	fmt.Printf("Starting server on port %v\n", portNumber)
	http.ListenAndServe(portNumber, nil)
}
