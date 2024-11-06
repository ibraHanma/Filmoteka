package controller

import (
	service2 "Filmoteka/internal/service"
)

type filmoteka struct {
	movie serviceMovie
	actor serviceActor
}

func NewFilmoteka(MovieService service2.MovieService, ActorService service2.ActorService) *filmoteka {
	return &filmoteka{movie: &MovieService, actor: &ActorService}
}
