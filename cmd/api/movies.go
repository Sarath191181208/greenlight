package main

import (
	"fmt"
	"net/http"
)

func (app *Application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Create new movie\n")
}

func (app *Application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	// getting the id param
	id, err := app.readIDParams(r)

	// if id param isn't found or is -ve
	if err != nil || id < 0 {
		http.NotFound(w, r)
		return
	}

	// returning the response
	fmt.Fprintf(w, "Show the details of the movie %d\n", id)
}
