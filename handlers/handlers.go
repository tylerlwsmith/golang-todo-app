package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	"golang-todo-app/assets"
	"golang-todo-app/models"
	"golang-todo-app/repositories"
)

var indexTmpl = assets.MakeTmpl("templates/pages/index.tmpl")
var editTmpl = assets.MakeTmpl("templates/pages/edit.tmpl")
var notFoundTmpl = assets.MakeTmpl("templates/pages/404.tmpl")

func RedirectToIndex(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/tasks", http.StatusSeeOther)
}

func Index(w http.ResponseWriter, r *http.Request) {
	indexTmpl.Render(w, r, models.Page{
		Title:    "Hello, world!",
		Content:  "I am a page.",
		PageData: repositories.GetTasks(),
	})
}

func Store(w http.ResponseWriter, r *http.Request) {
	repositories.StoreTask(models.Task{
		Id:          int(time.Now().Unix()),
		Description: r.PostFormValue("description"),
		Completed:   false,
	})
	http.Redirect(w, r, "/tasks", http.StatusFound)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(404)
		notFoundTmpl.Render(w, r, models.Page{
			Title: "Not Found",
		})
		return
	}

	todo, err := repositories.GetTask(id)

	if err != nil {
		w.WriteHeader(404)
		notFoundTmpl.Render(w, r, models.Page{
			Title: "Not Found",
		})
		return
	}

	editTmpl.Render(w, r, models.Page{
		Title:    "Edit todo",
		Content:  "Edit the todo below.",
		PageData: todo,
	})
}

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		print(id)
	}

	http.Redirect(w, r, "/tasks", http.StatusSeeOther)

}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		return
	}

	repositories.DeleteTask(id)
	http.Redirect(w, r, "/tasks", http.StatusFound)
}
