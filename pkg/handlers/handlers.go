package handlers

import (
	"net/http"

	"github.com/bagashiz/Go-Basic-Web-App/pkg/config"
	"github.com/bagashiz/Go-Basic-Web-App/pkg/models"
	"github.com/bagashiz/Go-Basic-Web-App/pkg/render"
)

// Repo is a variable that holds the repository used by the handlers
var Repo *Repository

// Repository is a struct that holds the application configuration
type Repository struct {
	App *config.AppConfig
}

// NewRepo is a function that creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers is a function that sets repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

//* Handler Functions
// Home is the home page handler function
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// call the renderTemplate function to render the home page
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

// About is the about page handler function
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello World!"

	// call the renderTemplate function to render the about page
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
