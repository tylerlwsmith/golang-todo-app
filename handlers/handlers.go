package handlers

import (
	"html/template"
	"net/http"
	"time"

	"strconv"

	"github.com/gorilla/mux"

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

func Store(w http.ResponseWriter, r *http.Request) {
	repositories.StoreTodo(models.Todo{
		Id:          int(time.Now().Unix()),
		Description: r.PostFormValue("new-todo-description"),
		Completed:   false,
	})
	http.Redirect(w, r, "/todos", http.StatusFound)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		return
	}

	repositories.DeleteTodo(id)
	http.Redirect(w, r, "/todos", http.StatusFound)
}
