package main

import (
	"Filmoteka/internal/controller"
	"Filmoteka/internal/router"
	"log"
)

func main() {

	actorController := controller.ActorController{}
	movieController := controller.MovieController{}

	srv := router.NewServer(actorController, movieController)

	if err := srv.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
