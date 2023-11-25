package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/todos", http.StatusSeeOther)
	}).Methods("GET")

	r.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>Hello, world!</h1>")
	}).Methods("GET")

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
