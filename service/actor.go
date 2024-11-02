package service

import (
	"errors"
	"time"
)

type storeActor interface {
	CreateActor(actor Actor) error
	GetActor(id int) (Actor, error)
	UpdateActor(id int, actor Actor) error
	DeleteActor(id int) error
}
type Actor struct {
	ID       int
	Name     string
	Birthday time.Time
	Gender   string
}

type ActorService struct {
	store storeActor
}

func NewActor(store storeActor) ActorService {
	return ActorService{store: store}
}

func (a *ActorService) CreateActor(actor Actor) error {
	if actor.Name == "" {
		return errors.New("actor name cannot be empty")
	}
	return a.store.CreateActor(actor)
}

// GetActor возвращает актера по ID
func (a *ActorService) GetActor(id int) (Actor, error) {
	if id <= 0 {
		return Actor{}, errors.New("actor id should be greater than  zero")
	}
	return a.store.GetActor(id)
}

// UpdateActor обновляет данные существующего актера
func (a *ActorService) UpdateActor(id int, actor Actor) error {
	if actor.Name == "" {
		return errors.New("actor name cannot be empty")
	}
	return a.store.UpdateActor(id, actor)
}

// DeleteActor удаляет актера по ID
func (a *ActorService) DeleteActor(id int) error {
	if id <= 0 {
		return errors.New("actor id should be greater than  zero")
	}
	return a.store.DeleteActor(id)
}
