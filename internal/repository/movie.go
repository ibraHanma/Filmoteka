package repository

import (
	"database/sql"
)

type Movie struct {
	ID          int
	Title       string
	Description string
	ReleaseDate int
	Rating      int
}

type movie struct {
	db *sql.DB
}

func NewMovie(db *sql.DB) *movie {
	return &movie{db: db}

}



func (m *movie) CreateMovie(movie Movie) (int, error) {
	var id int
	query := `INSERT INTO movie (title, description, release_date, rating) VALUES ($1, $2, $3, $4) RETURNING id`
	err := m.db.QueryRow(query, movie.Title, movie.Description, movie.ReleaseDate, movie.Rating).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}



func (m *movie) GetMovie(id int) (Movie, error) {
	var movie Movie
	query := `SELECT id,title,description,release_date,rating FROM movie WHERE id = $1`
	err := m.db.QueryRow(query, id).Scan(&movie.ID, &movie.Title, &movie.Description, &movie.ReleaseDate, &movie.Rating)
	if err != nil {
		return Movie{}, err
	}
	return movie, nil
}

func (m *movie) UpdateMovie(movie Movie) error {

	query := `UPDATE movie SET title = $1, description = $2, release_date = $3, rating = $4 WHERE id = $5`
	_, err := m.db.Exec(query, movie.Title, movie.Description, movie.ReleaseDate, movie.Rating, movie.ID)
	if err != nil {
		return err
	}
	return nil
}
func (m *movie) DeleteMovie(id int) error {
	query := `DELETE  FROM movie WHERE id = $1`
	_, err := m.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil

}
