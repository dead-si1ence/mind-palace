package main

import (
	"log"
	"net/http"
	"github.com/dead-si1ence/mind-palace/handlers"
)

func main() {
	// Serve static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Set up routes
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/about", handlers.AboutHandler)
	http.HandleFunc("/contact", handlers.ContactHandler)
	// Add more routes as needed

	// Start the server
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}