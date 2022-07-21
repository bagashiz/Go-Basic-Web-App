package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// RenderTemplate is a function that renders the specified template to the response writer
func RenderTemplate(w http.ResponseWriter, html string) {
	// Load the template from the templates folder and parse it into a html template object
	parsedTemplate, _ := template.ParseFiles("./templates/" + html)
	err := parsedTemplate.Execute(w, nil)
	// check for any errors
	if err != nil {
		fmt.Printf("Error parsing template: %v\n", err)
	}
}
