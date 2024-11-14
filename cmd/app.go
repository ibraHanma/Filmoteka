package main

import (
	"Filmoteka/internal/controller"
	"Filmoteka/internal/postgres"
	"Filmoteka/internal/repository"
	service2 "Filmoteka/internal/service"
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

	movieRepo := repository.NewMovie(db)
	actorRepo := repository.NewActor(db)

	movieService := service2.NewMovie(movieRepo)
	actorService := service2.NewActor(actorRepo)

	// Инициализация контроллера
	filmotekaController := controller.NewFilmoteka(movieService, actorService)

	return nil
}
