package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// NoSurf adds the CSRF (Cross Site Request Forgery)  to all POST requests.

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   app.IsProduction,
	})

	return csrfHandler
}

// SessionLoad loads and saves the session every request
func SessionLoad(next http.Handler) http.Handler {

	return session.LoadAndSave(next)
}
