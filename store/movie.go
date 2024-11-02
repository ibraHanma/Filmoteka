package store

import (
	"Filmoteka/service"
	"database/sql"
	"fmt"
)

type movie struct {
	db *sql.DB
}

func NewMovie(db *sql.DB) *movie {
	return &movie{db: db}
}

type MovieStore struct {
	movie []service.Movie
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
