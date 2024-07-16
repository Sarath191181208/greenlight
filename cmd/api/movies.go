package main

import (
	"net/http"
	"time"

	"sarath/greenlight/internal/data"
	"sarath/greenlight/internal/validator"
)

func (app *Application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
  // decoding into an intermediary struct  to handle 
  // maliciouse parameter passing of fields like 
  // `id` (or) `version` etc.
	var input struct {
		Title   string       `json:"title"`
		Year    int32        `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Genres  []string     `json:"genres"`
	}

  // check if valid json data is passed
	err := app.readJSON(&input, w, r)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	movie := data.Movie{
		Title:   input.Title,
		Generes: input.Genres,
		Year:    input.Year,
		Runtime: input.Runtime,
	}

	// validate the json data send Errors on invalid request
	v := validator.New()
	if data.ValidateMovie(v, movie); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
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
