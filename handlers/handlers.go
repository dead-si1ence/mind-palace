package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var templates = template.Must(template.ParseGlob(filepath.Join("templates", "*.html")))

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html", nil)
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html", nil)
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "contact.html", nil)
}

// Add more handlers as needed