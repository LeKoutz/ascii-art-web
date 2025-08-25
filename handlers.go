package handlers

import (
	"ascii-art-web/services"
	"html/template"
	"net/http"
)

type AsciiHandler struct {
	service *services.AsciiArtWeb
}

// PageData is used for passing data to the html file templates
type PageData struct {
	InputText   string
	InputBanner string
	AsciiArt    string
}

// ErrorData is used for passing data to the html file error templates
type ErrorData struct {
	StatusCode int
	Message    string
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

// HandleAsciiArt processes form submissions and returns ASCII art
func (a *AsciiHandler) HandleAsciiArt(w http.ResponseWriter, r *http.Request) {
	// Only allow POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Parse the form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	// Get the text and banner style from the form
	text := r.FormValue("text")
	if text == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	banner := r.FormValue("banner")
	if banner == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	// ===========================================
	// We should add a ValidateInput function here
	// ===========================================

	// Generate ASCII art
	asciiArt, err := a.service.Generate(text, banner)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Return the ASCII art
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(asciiArt))
}
