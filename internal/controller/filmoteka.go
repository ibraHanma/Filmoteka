package controller

import (
	service2 "Filmoteka/internal/service"
)

type Filmoteka struct {
	movie service2.MovieService
	actor service2.ActorService
}

func NewFilmoteka(movieService *service2.MovieService, actorService *service2.ActorService) *Filmoteka {
	return &Filmoteka{movie: *movieService, actor: *actorService}
}
