package main

import (
	"Filmoteka/internal/postgres"
	"Filmoteka/internal/repository"
	"Filmoteka/service"
	"fmt"
)

type movieStore struct {
	movie []service.Movie
}

func (s *movieStore) CreateMovie(movie service.Movie) error {
	s.movie = append(s.movie, movie)
	return nil
}

func (s *movieStore) GetMovie(id int) (*service.Movie, error) {
	for _, movie := range s.movie {
		if movie.ID == id {
			return &movie, nil
		}
	}
	return nil, fmt.Errorf("фильм не найден")
}
func (s *movieStore) UpdateMovie(movie service.Movie) error {
	for i, m := range s.movie {
		if m.ID == movie.ID {
			s.movie[i] = movie
			return nil
		}
	}
	return fmt.Errorf("фильм не найден")
}
func (s *movieStore) DeleteMovie(id int) error {
	for i, m := range s.movie {
		if m.ID == id {
			s.movie = append(s.movie[:i], s.movie[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("фильм не найден")
}

type actorStore struct {
	actor []service.Actor
}

func (r *actorStore) CreateActor(actor service.Actor) error {
	r.actor = append(r.actor, actor)
	return nil

}
func (r *actorStore) GetActor(id int) (*service.Actor, error) {
	for _, actor := range r.actor {
		if actor.ID == id {
			return &actor, nil
		}
	}
	return nil, fmt.Errorf("актер не найден")

}
func (r *actorStore) UpdateActor(actor *service.Actor) error {
	for i, m := range r.actor {
		if m.ID == actor.ID {
			r.actor[i] = *actor
			return nil
		}
	}
	return fmt.Errorf("актер не найден")

}
func (r *actorStore) DeleteActor(id int) error {
	for i, m := range r.actor {
		if m.ID == id {
			r.actor = append(r.actor[:i], r.actor[i+1:]...)
		}
	}
	return fmt.Errorf("актер не найден")
}

func Run() error {

	db, err := postgres.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	repository.NewMovie(db)
	repository.NewActor(db)

	return nil

}
