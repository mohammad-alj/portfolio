package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func handlerFuncBuilder(tmplName string, data any) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles(
			"templates/layout.html",
			fmt.Sprintf("templates/%s.html", tmplName), // here is the to be rendered template
		)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		tmpl.Execute(w, data)
	}
}

var Index = handlerFuncBuilder("index", nil)
