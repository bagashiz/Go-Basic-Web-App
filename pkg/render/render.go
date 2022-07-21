package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// tc is a map that stores the templates that have been loaded
var tc = make(map[string]*template.Template)

// RenderTemplate is a function that renders the specified template to the response writer
func RenderTemplate(w http.ResponseWriter, h string) {
	// create variables to store the template and any errors
	var html *template.Template
	var err error

	// check if the template has already been loaded into the map
	_, inMap := tc[h]
	if !inMap {
		// create the template
		log.Printf("Creating template cache and adding %s to the cache\n", h)
		err = createTemplateCache(h)
		// check for any errors
		if err != nil {
			log.Printf("Error creating template cache: %v\n", err)

		} else {
			// get the template from the map
			log.Printf("Using cached template for %v\n", h)
		}

		// assign the template to the html variable and check for any errors
		html = tc[h]
		err = html.Execute(w, nil)
		// check for any errors
		if err != nil {
			log.Printf("Error executing template: %v\n", err)
		}
	}
}

func createTemplateCache(h string) error {
	// create the templates slice to store the templates that are loaded
	templates := []string{
		fmt.Sprintf("./templates/%s", h),
		"./templates/base.html",
	}

	// parse the templates
	html, err := template.ParseFiles(templates...)
	// check for any errors
	if err != nil {
		return err
	}

	// add the template to the map
	tc[h] = html

	return nil
}
