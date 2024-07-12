package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Delcaring the version global constant
const version = "1.0.0"

type Config struct {
	port int
	env  string
}

type Application struct {
	config Config
	logger log.Logger
}

func main() {
	var config Config

	// Reading the flags of the application
	flag.IntVar(&config.port, "port", 4000, "API server port")
	flag.StringVar(&config.env, "env", "dev", "Environment (dev | stag | production)")
	flag.Parse()

	// defining the application
	app := &Application{
		config: config,
		logger: *log.New(os.Stdout, "", log.Ldate|log.Ltime),
	}

	// creating the http server
	mux := http.NewServeMux()

	// routing
	mux.HandleFunc("/v1/healthcheck", app.healthCheckHandler)

	// putting in sensible defaults
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Minute,
		WriteTimeout: 10 * time.Minute,
	}

	// start the server
	app.logger.Printf("Starting %s server on %s", app.config.env, server.Addr)
	err := server.ListenAndServe()
	app.logger.Fatal(err)
}
