package controller

import (
	"time"
)

type serviceMovie interface {
	CreateMovie(title string, description string, releaseDate time.Time, rating int) (int, error)
	GetMovie(id int) (Movie, error)
	UpdateMovie(id int, title string, description string, releaseDate time.Time, rating int) error
	DeleteMovie(id int) error
}

type Movie struct {
	ID          int
	Title       string
	Description string
	ReleaseDate time.Time
	Rating      int
}
