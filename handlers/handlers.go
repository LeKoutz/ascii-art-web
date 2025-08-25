package handlers

import (
	"ascii-art-web/services"
	"fmt"
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
		a.HandleErrors(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}
	// Check if we are at the root path
	if r.URL.Path != "/" {
		a.HandleErrors(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}
	// Parse the home template
	template, err := template.ParseFiles("templates/index.html")
	if err != nil {
		a.HandleErrors(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	// Execute the home template
	err = template.Execute(w, nil)
	if err != nil {
		a.HandleErrors(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
}

// HandleAsciiArt processes form submissions and returns ASCII art
func (a *AsciiHandler) HandleAsciiArt(w http.ResponseWriter, r *http.Request) {
	// Only allow POST requests
	if r.Method != http.MethodPost {
		a.HandleErrors(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}
	// Parse the form data
	if err := r.ParseForm(); err != nil {
		a.HandleErrors(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	// Get the text and banner style from the form
	text := r.FormValue("text")
	if text == "" {
		a.HandleErrors(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	banner := r.FormValue("banner")
	if banner == "" {
		a.HandleErrors(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	// ===========================================
	// We should add a ValidateInput function here
	// ===========================================

	// Generate ASCII art
	asciiArt, err := a.service.Generate(text, banner)
	if err != nil {
		a.HandleErrors(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	// Parse the home template
	template, err := template.ParseFiles("templates/index.html")
	if err != nil {
		a.HandleErrors(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	// Execute the home template
	err = template.Execute(w, PageData{InputText: text, InputBanner: banner, AsciiArt: asciiArt})
	if err != nil {
		a.HandleErrors(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
}

// HandleErrors serves error pages based on status codes
func (a *AsciiHandler) HandleErrors(w http.ResponseWriter, statusCode int, message string) {
	// Parse the error template
	template, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error %d: %s", statusCode, http.StatusText(statusCode)), statusCode)
		return
	}
	// Execute the error template
	err = template.Execute(w, ErrorData{StatusCode: statusCode, Message: message})
	if err != nil {
		http.Error(w, fmt.Sprintf("Error %d: %s", statusCode, http.StatusText(statusCode)), statusCode)
		return
	}
}
