package main

import (
	"fmt"
	"net/http"

	"github.com/mohammad-alj/portfolio/internal/handlers"
)

func main() {
	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	staticHandler := http.StripPrefix("/static/", fs)
	http.Handle("/static/", staticHandler)

	// Routes
	http.HandleFunc("/", handlers.Index)

	// Serve the website
	const PORT int = 3000
	fmt.Printf("Server running on http://localhost:%d\n", PORT)
	err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
	if err != nil {
		panic(err)
	}
}
