package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/ugandhar84/bookings/internal/config"
	"github.com/ugandhar84/bookings/internal/render"
	"log"
	"net/http"

	"github.com/ugandhar84/bookings/models"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandler(r *Repository) {
	Repo = r
}

// Home is the About page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(w, r, "contacts.page.tmpl", &models.TemplateData{})
}

// About is the About page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["test"] = "Hello String"
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{StringMap: stringMap})
}

func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["test"] = "Hello String"
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, r, "generals.page.tmpl", &models.TemplateData{StringMap: stringMap})
}

func (m *Repository) Suits(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["test"] = "Hello String"
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, r, "suits.page.tmpl", &models.TemplateData{StringMap: stringMap})
}

func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["test"] = "Hello String"
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{StringMap: stringMap})
}

// PostAvailability handler to handle POST message
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {

	start := r.Form.Get("start")
	end := r.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("start date is: %s  end date is: %s", start, end)))

}

type jsonResponse struct {
	OK      bool   `json:"ok""`
	Message string `json:"message"`
}

// AvailabilityJSON send Json response to availability search
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK: true, Message: "Available!",
	}

	out, err := json.MarshalIndent(resp, "", "   ")

	if err != nil {
		log.Println(err)
	}
	log.Println(string(out))
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)

}

func (m *Repository) Reservations(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["test"] = "Hello String"
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, r, "reservations.page.tmpl", &models.TemplateData{StringMap: stringMap})
}
