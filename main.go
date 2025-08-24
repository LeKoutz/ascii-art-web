package main

import (
	"log"
	"net/http"
	"ascii-art-web/handlers"
	"ascii-art-web/services"
)

func main() {
	// Create service
	asciiService := services.NewAsciiArtWeb()
	if err := asciiService.LoadBanners(); err != nil {
		log.Fatalf("Failed to load banners: %v", err)
	}

	// Create handler
	asciiHandler := handlers.NewAsciiHandler(asciiService)

	// Setup only the two required routes
	http.HandleFunc("/", asciiHandler.HandleHome)
	http.HandleFunc("/ascii-art", asciiHandler.HandleAsciiArt)

	// Start server
	port := ":8080"
	log.Printf("Server starting on http://localhost%s", port)
	log.Println("Press Ctrl+C to stop the server")
	log.Fatal(http.ListenAndServe(port, nil))
}