package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

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
	r.HandleFunc("/tasks/{id}", handlers.Update).Methods("PUT")
	r.HandleFunc("/tasks/{id}", handlers.Delete).Methods("DELETE")

	p := os.Getenv("PORT")
	if p == "" {
		p = "3000"
	}

	// https://github.com/gorilla/mux?tab=readme-ov-file#graceful-shutdown
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	// Decouple listen and serve so I can display message _after_ listening
	// on an avialable port.
	// https://stackoverflow.com/a/53332769/7759523
	l, err := net.Listen("tcp", ":"+p)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	// https://github.com/gorilla/mux?tab=readme-ov-file#graceful-shutdown
	s := &http.Server{
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      middleware.UseFormMethod(r), // Pass our instance of gorilla/mux in.
	}

	fmt.Printf("Listening on port %v\n", p)

	go func() {
		if err := s.Serve(l); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	s.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
