package main

import (
	"Filmoteka/controller"
	"Filmoteka/internal/postgres"
	"Filmoteka/internal/repository"
	"Filmoteka/service"
	"Filmoteka/store"
)

func Run() error {

	db, err := postgres.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	repository.NewMovie(db)
	repository.NewActor(db)

	movieStore := store.NewMovie(db)
	actorStore := store.NewActor(db)

	MovieService := service.NewMovie(movieStore)
	ActorService := service.NewActor(actorStore)

	filmotekaController := controller.NewFilmoteka(MovieService, ActorService)

	return nil
}
