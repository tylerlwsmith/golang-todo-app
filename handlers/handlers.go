package handlers

import (
	"html/template"
	"net/http"

	"golang-todo-app/assets"
	"golang-todo-app/models"
	"golang-todo-app/repositories"
)

var tmplDirs = []string{
	"templates/layouts/*.tmpl",
	"templates/partials/*.tmpl",
	"templates/pages/*.tmpl",
}
var tmpl = template.Must(template.ParseFS(assets.TmplFiles, tmplDirs...))

func RedirectToIndex(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/todos", http.StatusSeeOther)
}

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "layout.tmpl", models.Page{
		Title:    "Hello, world!",
		Content:  "I am a page.",
		PageData: repositories.GetTodos(),
	})
}
