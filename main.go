package main

import (
	"fmt"
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

	http.ListenAndServe(":"+p, r)
}
