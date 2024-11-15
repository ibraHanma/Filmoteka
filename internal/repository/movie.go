package repository

import (
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"time"
)

type Movie struct {
	ID          int
	Title       string
	Description string
	ReleaseDate time.Time
	Rating      int
}

type MovieRepo struct {
	db *sql.DB
}

func NewMovie(db *sql.DB) *MovieRepo {
	return &MovieRepo{db: db}
}

func (m *MovieRepo) CreateMovie(movie *Movie) (int, error) {
	var id int

	query, args, err := squirrel.Insert("movie").
		Columns("title", "description", "release_date", "rating").
		Values(movie.Title, movie.Description, movie.ReleaseDate, movie.Rating).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return 0, err
	}

	err = m.db.QueryRow(query, args...).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("не удалось выполнить запрос %w", err)
	}
	return id, nil
}

func (m *MovieRepo) GetMovie(id int) (Movie, error) {
	var movie Movie
	query, args, err := squirrel.Select("id", "title", "description", "release_date", "rating").
		From("movie").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return Movie{}, err
	}
	err = m.db.QueryRow(query, args...).Scan(&movie.ID, &movie.Title, &movie.Description, &movie.ReleaseDate, &movie.Rating)
	if err != nil {
		return Movie{}, fmt.Errorf("не удалось отсканировать фильм с идентификатором%d,%w", id, err)
	}
	return movie, nil

}

func (m *MovieRepo) UpdateMovie(movie Movie) error {
	query, args, err := squirrel.Update("movie").
		Set("title", movie.Title).
		Set("description", movie.Description).
		Set("release_date", movie.ReleaseDate).
		Set("rating", movie.Rating).
		Where(squirrel.Eq{"id": movie.ID}).
		ToSql()
	if err != nil {
		return fmt.Errorf("не удалось создать SQL запрос: %w", err)
	}
	_, err = m.db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("не удалось обновить фильм с идентификатором %d: %w", movie.ID, err)
	}
	return nil
}

func (m *MovieRepo) DeleteMovie(id int) error {
	query, args, err := squirrel.Delete("movie").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return fmt.Errorf("не удалось создать SQL запрос: %w", err)
	}
	_, err = m.db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("не удалось удалить фильм с идентификатором %d: %w", id, err)
	}
	return nil

}
