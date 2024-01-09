package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/TezzBhandari/lecture-03/handlers"
)

func main() {

	// logger
	// you can log to files but here i'm just logging to os.Stdout which is terminal
	logger := log.New(os.Stdout, "product-api ", log.LstdFlags)

	// creates the handler
	// hh := handlers.NewHello(logger)
	ph := handlers.NewProduct(logger)

	// create new server mux and register the handlers
	sm := http.NewServeMux()
	// sm.Handle("/", hh)
	sm.Handle("/products", ph)

	// we need to configure timeouts in the server. security against denial of service attack where client start the request and just wait. if multiple client does that then the servers resources will ultimately full and no futher request are gonna be processed
	s := &http.Server{
		Addr:         ":9090",           // configure the bind address
		Handler:      sm,                // set the default handler
		ErrorLog:     logger,            // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// starts a new http server on different go routine so that it doesn't block gracefull shutdown
	// creates a new server
	go func() {
		logger.Println("starting server on port 9090")
		err := s.ListenAndServe()

		if err != nil {
			logger.Printf("Error starting the server %s\n", err)
			os.Exit(1)
		}
	}()

	// traps  sigterm and interrup signal and gracefully shutdowns the  server
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	// signal.Notify(sigChan, os.Kill)

	// blocks until a signal is receieved
	sig := <-sigChan

	logger.Println("Recevied terminate, graceful shutdown", sig)

	// , waiting max 30 seconds for current operations to complete
	timeOutCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	// this  frees up resources
	// always call it
	defer cancel()
	s.Shutdown(timeOutCtx)

}
