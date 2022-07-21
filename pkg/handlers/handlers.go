package handlers

import (
	"net/http"

	"github.com/bagashiz/Go-Basic-Web-App/pkg/render"
)

//* Handler Functions
// Home is the home page handler function
func Home(w http.ResponseWriter, r *http.Request) {
	// call the renderTemplate function to render the home page
	render.RenderTemplate(w, "home.page.html")
}

// About is the about page handler function
func About(w http.ResponseWriter, r *http.Request) {
	// call the renderTemplate function to render the about page
	render.RenderTemplate(w, "about.page.html")
}
