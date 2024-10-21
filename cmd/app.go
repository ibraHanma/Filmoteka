package main

import (
	"Filmoteka/internal/postgres"
	"Filmoteka/internal/repository"
)

// запуск приложения

func Run() error {

	db, err := postgres.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	repository.NewMovie(db)
	repository.NewActor(db)

	return nil

}
