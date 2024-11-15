package store

import (
	"Filmoteka/internal/repository"
	"Filmoteka/internal/service"
)

type Movie struct {
	repo *repository.MovieRepo // или как вы назвали свой репозиторий
}

func (m Movie) CreateMovie(movie service.Movie) (int, error) {

	//TODO implement me
	panic("implement me")
}

func (m Movie) GetMovie(id int) (service.Movie, error) {
	//TODO implement me
	panic("implement me")
}

func (m Movie) UpdateMovie(id int, movie service.Movie) error {
	//TODO implement me
	panic("implement me")
}

func (m Movie) DeleteMovie(id int) error {
	//TODO implement me
	panic("implement me")
}

type MovieStore struct {
	movie []service.Movie
}

func NewMovie(repo *repository.MovieRepo) *Movie {
	return &Movie{repo: repo}
}

//
//func (s *MovieStore) CreateMovie(movie service.Movie) error {
//	s.movie = append(s.movie, movie)
//	return nil
//}
//
//func (s *MovieStore) GetMovie(id int) (*service.Movie, error) {
//	for _, movie := range s.movie {
//		if movie.ID == id {
//			return &movie, nil
//		}
//	}
//	return nil, fmt.Errorf("фильм не найден")
//}
//func (s *MovieStore) UpdateMovie(movie service.Movie) error {
//	for i, m := range s.movie {
//		if m.ID == movie.ID {
//			s.movie[i] = movie
//			return nil
//		}
//	}
//	return fmt.Errorf("фильм не найден")
//}
//func (s *MovieStore) DeleteMovie(id int) error {
//	for i, m := range s.movie {
//		if m.ID == id {
//			s.movie = append(s.movie[:i], s.movie[i+1:]...)
//			return nil
//		}
//	}
//	return fmt.Errorf("фильм не найден")
//}
