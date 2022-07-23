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

	// set UseCache to false
	app.UseCache = false

	// create a variable that holds the repository
	repo := handlers.NewRepo(&app)

	// pass the repository to the handlers
	handlers.NewHandlers(repo)

	// call the renderTemplate function to render the pages
	render.NewTemplates(&app)

	// create a variable to serve the application
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	// call the srv.ListenAndServe function to start the server
	err = srv.ListenAndServe()
	// check for any errors
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}

	fmt.Printf("Starting server on port %v\n", portNumber)
}
