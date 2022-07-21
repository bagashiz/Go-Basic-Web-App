package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/bagashiz/Go-Basic-Web-App/pkg/config"
	"github.com/bagashiz/Go-Basic-Web-App/pkg/models"
)

// functions is a variable that holds the FuncMap for the templates
var functions = template.FuncMap{}

// app is a variable that holds the application configuration
var app *config.AppConfig

// NewTemplates is a function that sets the application configuration to the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// AddDefaultFunctions is a function that adds the default data from TemplateData to the pages
func AddDefaultData(td *models.TemplateData) *models.TemplateData {

	return td
}

// RenderTemplate is a function that renders a template
func RenderTemplate(w http.ResponseWriter, html string, td *models.TemplateData) {
	// create a variable that holds the template cache
	var tc map[string]*template.Template

	// check if the cache is in use
	if app.UseCache {
		// get the template cache from the application configuration
		tc = app.TemplateCache
	} else {
		// create a new template cache
		tc, _ = CreateTemplateCache()
	}

	// get the requested template from the cache
	h, ok := tc[html]
	// check if the template is in the cache
	if !ok {
		// exit the application
		log.Fatal("Cannot get template: ", html)
	}

	// create buffer to hold the rendered html
	buf := new(bytes.Buffer)

	// set the template data
	td = AddDefaultData(td)

	// execute the template
	h.Execute(buf, td)

	// render the template
	_, err := buf.WriteTo(w)
	// check for any errors
	if err != nil {
		log.Printf("Error writing template %v: %v\n", html, err)
	}
}

// CreateTemplateCache is a function that creates a template cache to store the templates in memory for faster rendering
func CreateTemplateCache() (map[string]*template.Template, error) {
	// create a map to hold the templates
	myCache := map[string]*template.Template{}

	// get all of the html files in the templates directory
	pages, err := filepath.Glob("./templates/*.html")
	// check for any errors
	if err != nil {
		return myCache, err
	}

	// loop through the html files
	for _, page := range pages {
		// get the file name
		pageName := filepath.Base(page)

		// parse the html file into a template object
		templateSet, err := template.New(pageName).ParseFiles(page)
		// check for any errors
		if err != nil {
			return myCache, err
		}

		// check if the template is already in the cache
		matches, err := filepath.Glob("./templates/*.layout.html")
		// check for any errors
		if err != nil {
			return myCache, err
		}

		// check if there are any layout templates
		if len(matches) > 0 {
			// parse the layout template into a template object
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.html")
			// check for any errors
			if err != nil {
				return myCache, err
			}

			myCache[pageName] = templateSet
		}
	}
	return myCache, nil
}
