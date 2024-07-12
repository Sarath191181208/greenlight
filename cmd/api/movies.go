package main

import (
	"fmt"
	"net/http"
	"time"

	"sarath/greenlight/internal/data"
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

  err = app.writeJSON(movie, w, http.StatusOK, nil)

  if err != nil{
    http.Error(w, "The server encounterd a problem and could handle your request", http.StatusInternalServerError )
  }
}
