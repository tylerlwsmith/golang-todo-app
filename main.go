package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"golang-todo-app/handlers"
	"golang-todo-app/middleware"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.RedirectToIndex).Methods("GET")
	r.HandleFunc("/tasks", handlers.Index).Methods("GET")
	r.HandleFunc("/tasks", handlers.Store).Methods("POST")
	r.HandleFunc("/tasks/{id}", handlers.Edit).Methods("GET")
	r.HandleFunc("/tasks/{id}", handlers.Delete).Methods("DELETE")

	p := os.Getenv("PORT")
	if p == "" {
		p = "3000"
	}

	// Allow app to print listening message after starting the server.
	// https://stackoverflow.com/a/48250354/7759523
	done := make(chan bool)

	go http.ListenAndServe(":"+p, middleware.UseFormMethod(r))

	// Not using fmt.Printf because log.Printf works better with threads.
	// https://stackoverflow.com/a/41390023/7759523
	log.Printf("Now listening at http://localhost:%v", p)

	<-done
}
