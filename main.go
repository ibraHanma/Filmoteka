package main

import (
	"Filmoteka/internal/controller"
	"Filmoteka/internal/service"
	"Filmoteka/router"
)

func main() {

	actorService := service.ActorService{} // создание сервиса актера
	movieService := service.MovieService{} // создание сервиса фильма

	actorController := controller.ActorController{ActorService: actorService}
	movieController := controller.MovieController{MovieService: movieService}

	srv := router.NewServer(actorController, movieController)

	router.InitRoutes(&actorController, &movieController)

	// Запускаем сервер
	if err := srv.Run(); err != nil {
		panic(err)
	}
}
