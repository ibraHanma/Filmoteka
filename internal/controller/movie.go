package controller

import (
	"time"
)

type serviceMovie interface {
	CreateMovie(title string, description string, releaseDate time.Time, rating int) error
	GetMovie(id int) (string, string, int, error)
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
type MovieController struct {
	service serviceMovie
}

func (mc *MovieController) CreateMovie(title string, description string, releaseDate time.Time, rating int) error {
	return mc.service.CreateMovie(title, description, releaseDate, rating)

}
func (mc *MovieController) GetMovie(id int) (string, string, int, error) {
	return mc.service.GetMovie(id)

}

func (mc *MovieController) UpdateMovie(id int, title string, description string, releaseDate time.Time, rating int) error {
	return mc.service.UpdateMovie(id, title, description, releaseDate, rating)

}
func (mc *MovieController) DeleteMovie(id int) error {
	return mc.service.DeleteMovie(id)
}
