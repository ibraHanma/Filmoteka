package service

import "Filmoteka/controller"

type filmoteka struct {
	movie controller.ServiceMovie
	actor controller.ServiceActor
}

func NewFilmoteka(movieService controller.ServiceMovie, actorService controller.ServiceActor) *filmoteka {
	return &filmoteka{
		movie: movieService,
		actor: actorService,
	}
}
