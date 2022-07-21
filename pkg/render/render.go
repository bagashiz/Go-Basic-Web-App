package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// RenderTemplate is a function that renders a template
func RenderTemplate(w http.ResponseWriter, html string) {
	// create template cache
	templateCache, err := CreateTemplateCache()
	// check for any errors
	if err != nil {
		// exit the application
		log.Fatal(err)
	}

	// get the requested template from the cache
	h, ok := templateCache[html]
	// check if the template is in the cache
	if !ok {
		// exit the application
		log.Fatal(err)
	}

	// create buffer to hold the rendered html
	buf := new(bytes.Buffer)

	// execute the template
	h.Execute(buf, nil)

	// render the template
	_, err = buf.WriteTo(w)
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
