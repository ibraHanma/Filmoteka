package store

import (
	"Filmoteka/internal/service"
	"database/sql"
	"fmt"
)

type Movie struct {
	db *sql.DB
}

type MovieStore struct {
	movie []service.Movie
}

func NewMovie(db *sql.DB) *Movie {
	return &Movie{db: db}
}

func (s *MovieStore) CreateMovie(movie service.Movie) error {
	s.movie = append(s.movie, movie)
	return nil
}

func (s *MovieStore) GetMovie(id int) (*service.Movie, error) {
	for _, movie := range s.movie {
		if movie.ID == id {
			return &movie, nil
		}
	}
	return nil, fmt.Errorf("фильм не найден")
}
func (s *MovieStore) UpdateMovie(movie service.Movie) error {
	for i, m := range s.movie {
		if m.ID == movie.ID {
			s.movie[i] = movie
			return nil
		}
	}
	return fmt.Errorf("фильм не найден")
}
func (s *MovieStore) DeleteMovie(id int) error {
	for i, m := range s.movie {
		if m.ID == id {
			s.movie = append(s.movie[:i], s.movie[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("фильм не найден")
}
