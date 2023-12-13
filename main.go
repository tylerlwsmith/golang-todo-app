package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"golang-todo-app/handlers"
)

func useHttpFormMethod(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			m := r.PostFormValue("_method")
			if m != "" {
				r.Method = m
			}
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.RedirectToIndex).Methods("GET")
	r.HandleFunc("/todos", handlers.Index).Methods("GET")
	r.HandleFunc("/todos", handlers.Store).Methods("POST")
	r.HandleFunc("/todos/{id}", handlers.Delete).Methods("DELETE")

	p := os.Getenv("PORT")
	if p == "" {
		p = "3000"
	}

	// Allow app to print listening message after starting the server.
	// https://stackoverflow.com/a/48250354/7759523
	done := make(chan bool)

	go http.ListenAndServe(":"+p, useHttpFormMethod(r))

	// Not using fmt.Printf because log.Printf works better with threads.
	// https://stackoverflow.com/a/41390023/7759523
	log.Printf("Now listening at http://localhost:%v", p)

	<-done
}
