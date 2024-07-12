package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *Application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthCheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)

	return router
}