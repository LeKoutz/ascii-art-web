package handlers

import (
	"ascii-art-web/services"
	"html/template"
	"net/http"
)

type AsciiHandler struct {
	service *services.AsciiArtWeb
}

// NewAsciiHandler initializes a new AsciiHandler instance to be used by each request
func NewAsciiHandler(service *services.AsciiArtWeb) *AsciiHandler {
	return &AsciiHandler{
		service: service,
	}
}

// HandleHome serves the main page of the ASCII art web service
func (a *AsciiHandler) HandleHome(w http.ResponseWriter, r *http.Request) {
	// Only allow GET requests
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Check if we are at the root path
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// Parse the home template
	template, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Execute the home template
	template.Execute(w, nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
