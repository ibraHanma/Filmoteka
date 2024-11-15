package cmd

import (
	"Filmoteka/internal/controller"
	"Filmoteka/internal/postgres"
	"Filmoteka/internal/repository"
	"Filmoteka/internal/server"
	service2 "Filmoteka/internal/service"
	store2 "Filmoteka/internal/store"
	"log"
)

func Run() error {
	db, err := postgres.Connect()
	if err != nil {
		return err
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Ошибка при закрытии базы данных: %v", err)
		}
	}()

	movieRepo := repository.NewMovie(db)
	actorRepo := repository.NewActor(db)

	movieStore := store2.NewMovie(movieRepo)
	actorStore := store2.NewActor(actorRepo)

	movieService := service2.NewMovie(movieStore)
	actorService := service2.NewActor(actorStore)

	movieController := controller.MovieController{
		MovieService: movieService,
	}
	actorController := controller.ActorController{
		ActorService: actorService,
	}
	srv := server.NewServer(actorController, movieController)

	srv.InitRoutes()

	controller.NewFilmoteka(&movieService, &actorService)

	return nil
}
