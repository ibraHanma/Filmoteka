package model

import "time"

type Movie struct {
	ID          int
	Title       string
	Description string
	ReleaseDate time.Time
	Rating      int
}
