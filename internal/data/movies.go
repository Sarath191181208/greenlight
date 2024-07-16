package data

import (
	"time"

	"sarath/greenlight/internal/validator"
)

type Movie struct {
	ID        int64     `json:"id"`                // unique id of obj
	CreatedAt time.Time `json:"-"`                 // time at the movie is added to out db
	Title     string    `json:"title"`             // title of the movie
	Year      int32     `json:"year,omitempty"`    // Movie release year
	Runtime   Runtime   `json:"runtime,omitempty"` // Movie runtime in minutes
	Generes   []string  `json:"generes,omitempty"` // slice of generes of the movie
	Version   int32     `json:"version"`           // track number of updates to this particular movie record
}

func ValidateMovie(v *validator.Validator, movie Movie) {
	v.Check(movie.Title != "", "title", "must be provided")
	v.Check(len(movie.Title) <= 500, "title", "must not be more than 500 bytes long")

	v.Check(movie.Year != 0, "year", "must be provided and shouldn't be zero")
	v.Check(movie.Year >= 1888, "year", "must be greater than 1888")
	v.Check(movie.Year <= int32(time.Now().Year()), "year", "must not be in the future")

	v.Check(movie.Runtime != 0, "runtime", "must be provided")
	v.Check(movie.Runtime > 0, "runtime", "must be a positive integer")

	v.Check(movie.Generes != nil, "generes", "must be provided")
	v.Check(len(movie.Generes) >= 1, "generes", "must contain at least 1 genre")
	v.Check(len(movie.Generes) <= 5, "generes", "must not contain more than 5 generes")
	v.Check(validator.Unique(movie.Generes), "generes", "must not contain duplicate values")
}
