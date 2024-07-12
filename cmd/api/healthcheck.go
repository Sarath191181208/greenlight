package main

import (
	"fmt"
	"net/http"
)

func (app *Application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Status: Available\n")
  fmt.Fprintf(w, "environment %s\n", app.config.env)
  fmt.Fprintf(w, "version: %s\n", version)
}
