package handlers

import (
	"html/template"
	"net/http"

	"golang-todo-app/assets"
	"golang-todo-app/models"
	"golang-todo-app/repositories"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFS(assets.TmplFiles, "templates/*.html"))
}

func RedirectToIndex(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/todos", http.StatusSeeOther)
}

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.go.html", models.Page{
		Title:    "Hello, world!",
		Content:  "I am a page.",
		PageData: repositories.GetTodos(),
	})
}
