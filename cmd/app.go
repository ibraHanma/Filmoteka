package main

import (
	"Filmoteka/internal/controller"
	"Filmoteka/internal/postgres"
	"Filmoteka/internal/repository"
	service2 "Filmoteka/internal/service"
	store2 "Filmoteka/internal/store"
	"database/sql"
)

func Run() error {

	db, err := postgres.Connect()
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	repository.NewMovie(db)
	repository.NewActor(db)

	movieStore := store2.NewMovie(db)
	actorStore := store2.NewActor(db)

	MovieService := service2.NewMovie(movieStore)
	ActorService := service2.NewActor(actorStore)

	filmotekaController := controller.NewFilmoteka(MovieService, ActorService)

	return nil
}
