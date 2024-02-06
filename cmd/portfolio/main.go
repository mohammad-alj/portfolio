package main

import (
	"fmt"
	"net/http"

	"github.com/mohammad-alj/portfolio/internal/handlers"
)

func main() {
	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Routes
	http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/toggle-lang", handlers.ToggleLangHandler)

	// Serve the website
	const PORT int = 3000
	fmt.Printf("Server running on http://localhost:%d\n", PORT)
	err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
	if err != nil {
		panic(err)
	}
}
