package service

import "time"

// Интерфейс для операций с акторами
type StoreActor interface {
	CreateActor(actor Actor) error
	GetActor(id int) (Actor, error)
	UpdateActor(id int, actor Actor) error
	DeleteActor(id int) error
}

// Структура актора
type Actor struct {
	ID       int
	Name     string
	Birthday time.Time
	Gender   string
}
type actor struct {
	store StoreActor
}

// Конструктор для создания нового сервиса актеров
func NewActor(store StoreActor) actor {
	return actor{store: store}
}

// Методы сервиса актера (CRUD)
func (a *actor) CreateActor(actor Actor) error {
	return a.store.CreateActor(actor)

}

func (a *actor) GetActor(id int) (Actor, error) {
	return a.store.GetActor(id)
}

func (a *actor) UpdateActor(id int, actor Actor) error {
	return a.store.UpdateActor(id, actor)
}

func (a *actor) DeleteActor(id int) error {
	return a.store.DeleteActor(id)
}
