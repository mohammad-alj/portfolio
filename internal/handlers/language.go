package handlers

import (
	"net/http"

	"github.com/mohammad-alj/portfolio/internal/utils"
)

func ToggleLangHandler(w http.ResponseWriter, r *http.Request) {
	c := getLangCookie(w, r)

	// Togglin
	if c.Value == "ar" {
		c.Value = "en"
	} else {
		c.Value = "ar"
	}

	http.SetCookie(w, c)
}

func getLangCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
	langCookie, err := r.Cookie("lang")

	// lang cookie will point to a new cookie if it isn't or is invalid
	if err != nil || !(langCookie.Value == "ar" || langCookie.Value == "en") {
		langCookie = utils.NewLangCookie("lang", "en")
		http.SetCookie(w, langCookie)
	}

	return langCookie
}
