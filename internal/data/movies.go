package data

import "time"

type Movie struct {
	ID        int64     `json:"id"`                       // unique id of obj
	CreatedAt time.Time `json:"-"`                        // time at the movie is added to out db
	Title     string    `json:"title"`                    // title of the movie
	Year      int32     `json:"year,omitempty"`           // Movie release year
	Runtime   int       `json:"runtime,omitempty,string"` // Movie runtime in minutes
	Generes   []string  `json:"generes,omitempty"`        // slice of generes of the movie
	Version   int32     `json:"version"`                  // track number of updates to this particular movie record
}
