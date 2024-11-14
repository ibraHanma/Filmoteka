package controller

import (
	service2 "Filmoteka/internal/service"
)

type Filmoteka struct {
	movie service2.MovieService
	actor service2.ActorService
}

func NewFilmoteka(MovieService *service2.MovieService, ActorService *service2.ActorService) *Filmoteka {
	return &Filmoteka{movie: *MovieService, actor: *ActorService}
}
