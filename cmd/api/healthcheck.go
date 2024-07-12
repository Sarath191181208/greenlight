package main

import (
	"net/http"
)

func (app *Application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := envelope{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

  err := app.writeJSON(data, w, http.StatusOK, nil)
  if err != nil{
    app.logger.Fatal(err)
    app.serverErrorResponse(w, r, err)
  }
}
