package main

import (
	"fmt"
	"github.com/ugandhar84/bookings/internal/config"
	"github.com/ugandhar84/bookings/internal/handlers"
	"github.com/ugandhar84/bookings/internal/render"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

var PortNumber string = ":8090"
var app config.AppConfig
var session *scs.SessionManager

// main Application entry point
func main() {
	tc, err := render.CreateCacheTemplates()
	app.TemplateCache = tc

	if err != nil {
		fmt.Println(err)
	}

	app.UseCache = false
	app.IsProduction = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)
	render.NewTemplates(&app)

	session = scs.New()
	app.Session = session
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.IsProduction

	fmt.Println("Starting server on port number", PortNumber)
	//_ = http.ListenAndServe(PortNumber, nil)

	serv := &http.Server{
		Addr:    PortNumber,
		Handler: routes(&app),
	}

	err = serv.ListenAndServe()
	if err != nil {
		log.Fatal("Start failed", err)

	}
}
