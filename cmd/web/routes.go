package main

import (
	"net/http"

	"github.com/bagashiz/Go-Basic-Web-App/pkg/config"
	"github.com/bagashiz/Go-Basic-Web-App/pkg/handlers"
	"github.com/bmizerany/pat"
)

// routes is a function that creates a new router and registers the handlers
func routes(app *config.AppConfig) http.Handler {
	// create a variable as a multiplexer
	mux := pat.New()

	// set the routes for all the pages in the application
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
