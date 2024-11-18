package service

import (
	"Filmoteka/internal/model"
	"errors"
	"time"
)

type storeActor interface {
	CreateActor(actor Actor) (int, error)
	GetActor(id int) (Actor, error)
	UpdateActor(id int, actor Actor) error
	DeleteActor(id int) error
}

type Actor struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Birthday time.Time `json:"birthday"`
	Gender   string    `json:"gender"`
}
type ActorService struct {
	store storeActor
	actor map[int]model.Actor
}

func NewActor(store storeActor) ActorService {
	return ActorService{store: store}
}

func (as *ActorService) CreateActor(actor Actor) (int, error) {
	if actor.Name == "" {
		return 0, errors.New("имя актера не может быть пустым")
	}
	if actor.Birthday.IsZero() {
		return 0, errors.New("дата рождения не может быть пустой")
	}
	if actor.Gender == "" {
		return 0, errors.New("пол актера не может быть пустым")
	}
	id, err := as.store.CreateActor(actor)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (as *ActorService) GetActor(id int) (Actor, error) {
	if id <= 0 {
		return Actor{}, errors.New("actor id should be greater than zero")
	}
	return as.store.GetActor(id)
}

func (as *ActorService) UpdateActor(id int, actor Actor) error {
	if actor.Name == "" {
		return errors.New("the actor's name cannot be empty")
	}
	if actor.Birthday.IsZero() {
		return errors.New("еhe date of birth cannot be empty")
	}
	if actor.Gender == "" {
		return errors.New("the actor's gender cannot be empty")
	}
	if id <= 0 {
		return errors.New("the actor's ID must be greater than zero")
	}
	return as.store.UpdateActor(id, actor)
}

func (as *ActorService) DeleteActor(id int) error {
	if id <= 0 {
		return errors.New("actor id should be greater than zero")
	}
	return as.store.DeleteActor(id)
}
