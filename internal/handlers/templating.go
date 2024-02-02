package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

type pageData struct {
	Lang string
}

var Index = templateHandlerBuilder("index")

func templateHandlerBuilder(tmplName string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the language cookie
		langCookie := getLangCookie(w, r)

		// Parse the dynamic template with the layout
		tmpl, err := template.ParseFiles(
			"templates/layout.html",
			fmt.Sprintf("templates/%s.html", tmplName),
		)

		// Advanced error handeling
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		// Define the page data
		data := pageData{
			Lang: langCookie.Value,
		}

		// Compile the template with a 200
		w.WriteHeader(http.StatusOK)
		tmpl.Execute(w, data)
	}
}
