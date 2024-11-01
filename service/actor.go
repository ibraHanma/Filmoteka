package service

import (
	"time"
)

// Интерфейс для операций с актерами
type storeActor interface {
	CreateActor(actor Actor) error
	GetActor(id int) (Actor, error)
	UpdateActor(id int, actor Actor) error
	DeleteActor(id int) error
}

// Структура актера

type Actor struct {
	ID       int
	Name     string
	Birthday time.Time
	Gender   string
}
type ActorService struct {
	store storeActor
}

// Конструктор для создания нового сервиса актеров

func NewActor(store storeActor) ActorService {
	return ActorService{store: store}
}

// Методы сервиса актера (CRUD)

func (a *ActorService) CreateActor(actor Actor) error {
	return a.store.CreateActor(actor)

}

func (a *ActorService) GetActor(id int) (Actor, error) {
	return a.store.GetActor(id)
}

func (a *ActorService) UpdateActor(id int, actor Actor) error {
	return a.store.UpdateActor(id, actor)
}

func (a *ActorService) DeleteActor(id int) error {
	return a.store.DeleteActor(id)
}
