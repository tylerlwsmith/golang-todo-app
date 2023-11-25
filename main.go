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
		fmt.Fprint(w, "<h1>Hello, world!</h1>")
	})

	p := os.Getenv("PORT")
	if p == "" {
		p = "3000"
	}

	// Allow app to print listening message after starting the server.
	// https://stackoverflow.com/a/48250354/7759523
	done := make(chan bool)

	go http.ListenAndServe(":"+p, r)
	log.Printf("Now listening at http://localhost:%v", p)

	<-done
}
