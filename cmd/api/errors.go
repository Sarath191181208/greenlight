package main

import (
	"fmt"
	"net/http"
)

func (app *Application) logError(_ *http.Request, err error) {
	app.logger.Println(err)
}

func (app *Application) errorResponse(w http.ResponseWriter, _ *http.Request, status int, message interface{}) {
	// getting the error msg
	errMsg := envelope{"error": message}

	// writing the error msg
	err := app.writeJSON(errMsg, w, status, nil)
	// if the message writer fails log it
	if err != nil {
		app.logError(nil, err)
		w.WriteHeader(500)
	}
}

func (app *Application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)
	message := "the server encountered a problem and could not process your request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

func (app *Application) notFoundResponse(w http.ResponseWriter, _ *http.Request) {
	msg := "the requested resource couldn't be found"
	app.errorResponse(w, nil, http.StatusNotFound, msg)
}

func (app *Application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf("the method %s isn't supported for this resource", r.Method)
	app.errorResponse(w, nil, http.StatusNotFound, msg)
}

func (app *Application) failedValidationResponse (w http.ResponseWriter, r *http.Request, errors map[string]string){
  app.errorResponse(w, r, http.StatusUnprocessableEntity, errors)
}
