package store

import (
	"Filmoteka/internal/repository"
	"Filmoteka/internal/service"
	"log"
)

func NewMovie(repo *repository.MovieRepo) *Movie {
	return &Movie{repo: repo}
}

type Movie struct {
	repo *repository.MovieRepo
}

func (m *Movie) CreateMovie(movie service.Movie) (int, error) {
	repoMovie := repository.Movie{
		Title:       movie.Title,
		Description: movie.Description,
		ReleaseDate: movie.ReleaseDate,
		Rating:      movie.Rating,
	}
	id, err := m.repo.CreateMovie(&repoMovie)
	if err != nil {
		log.Printf("Ошибка при создании фильма: %v", err)
		return 0, err
	}
	return id, nil
}

func (m Movie) GetMovie(id int) (service.Movie, error) {
	movie, err := m.repo.GetMovie(id)
	if err != nil {
		log.Printf("Ошибка при получении фильма с ID %d: %v", id, err)
		return service.Movie{}, err
	}
	return service.Movie(movie), nil
}

func (m *Movie) UpdateMovie(id int, movie service.Movie) error {
	repoMovie := repository.Movie{
		ID:          id,
		Title:       movie.Title,
		Description: movie.Description,
		ReleaseDate: movie.ReleaseDate,
		Rating:      movie.Rating,
	}
	err := m.repo.UpdateMovie(&repoMovie)
	if err != nil {
		log.Printf("Ошибка при обновлении фильма с ID %d: %v", id, err)
		return err
	}
	return nil
}

func (m Movie) DeleteMovie(id int) error {
	err := m.repo.DeleteMovie(id)
	if err != nil {
		log.Printf("Ошибка при удалении фильма с ID %d: %v", id, err)
		return err
	}
	return nil
}
