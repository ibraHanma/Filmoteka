package controller

import "Filmoteka/service"

type filmoteka struct {
	movie serviceMovie
	actor ServiceActor
}

func NewFilmoteka(MovieService service.MovieService, ActorService service.ActorService) *filmoteka {
	return &filmoteka{movie: MovieService, actor: ActorService}
}
