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

func getLangCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
	langCookie, err := r.Cookie("lang")
	if err != nil || !(langCookie.Value == "ar" || langCookie.Value == "en") {
		langCookie = &http.Cookie{
			Name:   "lang",
			Value:  "en",
			MaxAge: 365 * 24 * 60 * 60,
			Path:   "/",
		}
		http.SetCookie(w, langCookie)
	}
	return langCookie
}

func templateHandlerBuilder(tmplName string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		langCookie := getLangCookie(w, r)
		tmpl, err := template.ParseFiles(
			"templates/layout.html",
			fmt.Sprintf("templates/%s.html", tmplName), // here is the to be rendered template
		)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		data := pageData{
			Lang: langCookie.Value,
		}
		tmpl.Execute(w, data)
	}
}

func ToggleLangHandler(w http.ResponseWriter, r *http.Request) {
	c := getLangCookie(w, r)
	var val string

	// Togglin
	if c.Value == "ar" {
		val = "en"
	} else {
		val = "ar"
	}

	http.SetCookie(w, &http.Cookie{
		Name:   "lang",
		Value:  val,
		MaxAge: 365 * 24 * 60 * 60,
		Path:   "/",
	})
}
