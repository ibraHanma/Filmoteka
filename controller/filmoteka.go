package controller

import "Filmoteka/service"

type filmoteka struct {
	movie serviceMovie
	actor ServiceActor
}

func NewFilmoteka(movieService service.MovieService, actorService service.ActorService) *filmoteka {
	return &filmoteka{
		movie: movieService,
		actor: actorService,
	}
}
