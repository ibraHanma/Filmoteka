package service

import "time"

type StoreMovie interface {
	CreateMovie(movie Movie) (int, error)
	GetMovie(id int) (Movie, error)
	UpdateMovie(id int, movie Movie) error
	DeleteMovie(id int) (int, error)
}

type Movie struct {
	ID          int
	Title       string
	Description string
	ReleaseDate time.Time
	Rating      int
}

// Структура сервиса для фильмов
type movie struct {
	store StoreMovie
}

// Конструктор для создания нового сервиса фильмов
func NewMovie(store StoreMovie) movie {
	return movie{store: store}

}

// Методы сервиса фильма (CRUD)
func (m *movie) CreateMovie(movie Movie) (int, error) {
	return m.store.CreateMovie(movie)
}
func (m *movie) GetMovie(id int) (Movie, error) {
	return m.store.GetMovie(id)
}
func (m *movie) UpdateMovie(id int, movie Movie) error {
	return m.store.UpdateMovie(id, movie)
}
func (m *movie) DeleteMovie(id int) (int, error) {
	return m.store.DeleteMovie(id)
}
