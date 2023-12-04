package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"golang-todo-app/handlers"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.RedirectToIndex).Methods("GET")
	r.HandleFunc("/todos", handlers.Index).Methods("GET")

	p := os.Getenv("PORT")
	if p == "" {
		p = "3000"
	}

	// Allow app to print listening message after starting the server.
	// https://stackoverflow.com/a/48250354/7759523
	done := make(chan bool)

	go http.ListenAndServe(":"+p, r)

	// Not using fmt.Printf because log.Printf works better with threads.
	// https://stackoverflow.com/a/41390023/7759523
	log.Printf("Now listening at http://localhost:%v", p)

	<-done
}
