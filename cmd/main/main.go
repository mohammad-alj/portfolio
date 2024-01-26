package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	// Serve static files
	staticHandler := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	http.Handle("/static/", staticHandler)

	// Home
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/layout.html", "templates/index.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
			return
		}
		tmpl.Execute(w, nil)
	})

	const PORT int = 3000
	fmt.Printf("Server running on http://localhost:%d\n", PORT)
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
}
