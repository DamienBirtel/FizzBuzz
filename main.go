package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/DamienBirtel/FizzBuzz/handlers"
)

// these could be taken from environment variables, but this will do for now
var (
	address      = "localhost:9090"
	readTimeout  = 5 * time.Second
	writeTimeout = 5 * time.Second
	idleYimeout  = 120 * time.Second
)

func createServer() *http.Server {

	fbh := handlers.NewFizzBuzzHandler()

	sm := http.NewServeMux()
	sm.Handle("/", fbh)

	srv := &http.Server{
		Addr:         address,
		Handler:      sm,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		IdleTimeout:  idleYimeout,
	}
	return srv
}

func main() {

	// we create our server
	srv := createServer()

	// we plan te gracefully shutdown the server by catching an interrupt signal
	idleConnsClosed := make(chan struct{})
	go func() {
		defer close(idleConnsClosed)
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		signal.Notify(sigint, os.Kill)
		<-sigint

		// We received an interrupt signal, shut down.
		err := srv.Shutdown(context.Background())
		if err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}
	}()

	log.Println("Starting fizzbuzz-api server...")

	err := srv.ListenAndServe()
	if err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
}
