package data

import "time"

type Movie struct {
	ID        int64     // unique id of obj
	CreatedAt time.Time // time at the movie is added to out db
	Title     string    // title of the movie
	Year      int32     // Movie release year
	Runtime   int       // Movie runtime in minutes
	Generes   []string  // slice of generes of the movie
	Version   int32     // track number of updates to this particular movie record
}
