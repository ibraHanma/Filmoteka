package cmd

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

	movieRepo := repository.NewMovie(db)
	actorRepo := repository.NewActor(db)

	movieStore := store2.NewMovie(movieRepo)
	actorStore := store2.NewActor(actorRepo)

	// Инициализация сервисов
	movieService := service2.NewMovie(movieStore)
	actorService := service2.NewActor(actorStore)

	// Инициализация контроллера
	filmotekaController := controller.NewFilmoteka(movieService, actorService)

	return nil
}
