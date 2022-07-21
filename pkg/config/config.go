package config

import (
	"html/template"
	"log"
)

// AppConfig is a struct that holds the application configuration
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}
