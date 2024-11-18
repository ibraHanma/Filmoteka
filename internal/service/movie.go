package service

import (
	"Filmoteka/internal/model"
	"errors"
	"time"
)

type storeMovie interface {
	CreateMovie(movie Movie) (int, error)
	GetMovie(id int) (Movie, error)
	UpdateMovie(id int, movie Movie) error
	DeleteMovie(id int) error
}

type Movie struct {
	ID          int       `json:"ID"`
	Title       string    `json:"Title,"`
	Description string    `json:"Description"`
	ReleaseDate time.Time `json:"ReleaseDate"`
	Rating      int       `json:"Rating"`
}

type MovieService struct {
	store storeMovie
	movie map[int]model.Movie
}

func NewMovie(store storeMovie) MovieService {
	return MovieService{store: store}

}

func (m *MovieService) CreateMovie(movie Movie) (int, error) {
	if movie.Title == "" {
		return 0, errors.New("movie title is required")
	}
	if movie.Description == "" {
		return 0, errors.New("movie description is required")
	}
	if movie.ReleaseDate.IsZero() {
		return 0, errors.New("movie release date is required")
	}
	return m.store.CreateMovie(movie)
}
func (m *MovieService) GetMovie(id int) (Movie, error) {
	if id <= 0 {
		return Movie{}, errors.New("invalid movie ID")
	}
	return m.store.GetMovie(id)
}

func (m *MovieService) UpdateMovie(id int, movie Movie) error {
	if movie.Title == "" {
		return errors.New("movie title is required")
	}
	if movie.Description == "" {
		return errors.New("movie description is required")
	}
	if movie.ReleaseDate.IsZero() {
		return errors.New("movie release date is required")
	}
	return m.store.UpdateMovie(id, movie)
}

func (m *MovieService) DeleteMovie(id int) error {
	if id <= 0 {
		return errors.New("invalid movie ID")
	}
	return m.store.DeleteMovie(id)
}
