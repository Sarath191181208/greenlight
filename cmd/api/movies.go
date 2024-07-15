package main

import (
	"net/http"
	"time"

	"sarath/greenlight/internal/data"
)

func (app *Application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
  // read the json data given by the user
  var movie data.Movie
  err := app.readJSON(&movie, w, r) 
  if err != nil{
    app.errorResponse(w, r, http.StatusBadRequest, err)
    return 
  }
}

func (app *Application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	// getting the id param
	id, err := app.readIDParams(r)
	// if id param isn't found or is -ve
	if err != nil || id < 0 {
		http.NotFound(w, r)
		return
	}

	// getting the movie data
	movie := data.Movie{
		ID:        0,
		CreatedAt: time.Now(),
		Title:     "Movie man",
		Year:      2024,
		Runtime:   61,
		Generes:   []string{"Action", "Romance", "Fighiting"},
		Version:   1,
	}

	err = app.writeJSON(envelope{"movie": movie}, w, http.StatusOK, nil)
	if err != nil {
    app.serverErrorResponse(w, r, err)
	}
}
