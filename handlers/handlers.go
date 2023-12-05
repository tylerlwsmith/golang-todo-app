package handlers

import (
	"html/template"
	"net/http"

	"golang-todo-app/assets"
)

type Page struct {
	Title   string
	Content string
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFS(assets.TmplFiles, "templates/*.html"))
}

func RedirectToIndex(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/todos", http.StatusSeeOther)
}

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", Page{
		Title:   "Hello, world!",
		Content: "I am a page.",
	})
}
