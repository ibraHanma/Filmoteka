package store

import (
	"Filmoteka/internal/repository"
	"Filmoteka/internal/service"
	"log"
)

func NewActor(repo *repository.ActorRepo) *Actor {
	return &Actor{repo: repo}
}

type Actor struct {
	repo *repository.ActorRepo
}

func (a Actor) CreateActor(actor service.Actor) (int, error) {
	id, err := a.repo.CreateActor(repository.Actor(actor))
	if err != nil {
		log.Printf("Ошибка при создании актера: %v", err)
		return 0, err
	}
	return id, nil
}

func (a Actor) GetActor(id int) (service.Actor, error) {
	actor, err := a.repo.GetActor(id)
	if err != nil {
		log.Printf("Ошибка при получении актера с ID %d: %v", id, err)
		return service.Actor{}, err
	}
	return service.Actor(actor), nil
}

func (a Actor) UpdateActor(id int, actor service.Actor) error {
	err := a.repo.UpdateActor(repository.Actor(actor))
	if err != nil {
		log.Printf("Ошибка при обновлении актера с ID %d: %v", id, err)
		return err
	}
	return nil
}

func (a Actor) DeleteActor(id int) error {
	err := a.repo.DeleteActor(id)
	if err != nil {
		log.Printf("Ошибка при удалении актера с ID %d: %v", id, err)
		return err
	}
	return nil
}
