package render

import (
	"bytes"
	"fmt"
	"github.com/justinas/nosurf"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/ugandhar84/bookings/models"
	"github.com/ugandhar84/bookings/pkg/config"
)

// TemplateData Holds data from templates

var app *config.AppConfig

// NewTemplates set the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.CSRFToken = nosurf.Token(r)

	return td
}

// RenderTemplate for rendering the templates from html
func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {
	fmt.Println(tmpl)
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateCacheTemplates()
		fmt.Println(tc)
	}

	// Get requested template
	t, ok := tc[tmpl]
	fmt.Println(t)
	if !ok {
		log.Fatal("could not get template from path")
	}
	buf := new(bytes.Buffer)

	td = AddDefaultData(td, r)
	_ = t.Execute(buf, td)

	// Render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateCacheTemplates() (map[string]*template.Template, error) {

	var myCache = map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*page.tmpl")

	if err != nil {
		fmt.Println(err)
		return myCache, err
	}

	// Iterate through all templates
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)

		if err != nil {

			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*layout.tmpl")

		if err != nil {

			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*layout.tmpl")
			if err != nil {

				return myCache, err
			}
		}

		myCache[name] = ts

	}

	return myCache, nil
}
