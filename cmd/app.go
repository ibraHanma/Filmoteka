package main

import (
	"Filmoteka/internal/postgres"
)

func Run() error {

	db, err := postgres.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	return nil
}
