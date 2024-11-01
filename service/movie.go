package service

import "time"

type storeMovie interface {
	CreateMovie(movie Movie) (int, error)
	GetMovie(id int) (Movie, error)
	UpdateMovie(id int, movie Movie) error
	DeleteMovie(id int) error
}

type Movie struct {
	ID          int
	Title       string
	Description string
	ReleaseDate time.Time
	Rating      int
}

// Структура сервиса для фильмов

type MovieService struct {
	store storeMovie
}

// Конструктор для создания нового сервиса фильмов

func NewMovie(store storeMovie) *MovieService {
	return &MovieService{store: store}

}

// Методы сервиса фильма (CRUD)

func (m *MovieService) CreateMovie(movie Movie) (int, error) {
	return m.store.CreateMovie(movie)
}
func (m *MovieService) GetMovie(id int) (Movie, error) {
	return m.store.GetMovie(id)
}
func (m *MovieService) UpdateMovie(id int, movie Movie) error {
	return m.store.UpdateMovie(id, movie)
}
func (m *MovieService) DeleteMovie(id int) error {
	return m.store.DeleteMovie(id)
}
