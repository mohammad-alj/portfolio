package utils

import "net/http"

func NewLangCookie(name string, val string) *http.Cookie {
	return &http.Cookie{
		Name:   "lang",
		Value:  val,
		MaxAge: 365 * 24 * 60 * 60, // about a year
		Path:   "/",                // it can be accessed anywhere
	}
}
