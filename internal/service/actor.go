package service

import (
	"Filmoteka/internal/model"
	"errors"
	"time"
)

// Определяем интерфейс для хранилища актера
type storeActor interface {
	CreateActor(actor Actor) (int, error)  // Позволяет создавать актеров
	GetActor(id int) (Actor, error)        // Получает актера по ID
	UpdateActor(id int, actor Actor) error // Обновляет данные актера
	DeleteActor(id int) error              // Удаляет актера по ID
}

// Структура актера

type Actor struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Birthday time.Time `json:"birthday"`
	Gender   string    `json:"gender"`
}

// Сервис актера

type ActorService struct {
	store  storeActor // Хранилище, которое реализует интерфейс storeActor
	actors map[int]model.Actor
}

// Создание нового сервиса актера

func NewActor(store storeActor) ActorService {
	return ActorService{store: store}
}

// Создает нового актера

func (as *ActorService) CreateActor(actor Actor) (int, error) {
	if actor.Name == "" {
		return 0, errors.New("actor name cannot be empty")
	}
	// Сохраняем актера в хранилище
	id, err := as.store.CreateActor(actor)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// Получает актера по ID

func (as *ActorService) GetActor(id int) (Actor, error) {
	if id <= 0 {
		return Actor{}, errors.New("actor id should be greater than zero")
	}
	return as.store.GetActor(id)
}

// Обновляет данные существующего актера

func (as *ActorService) UpdateActor(id int, actor Actor) error {
	if actor.Name == "" {
		return errors.New("actor name cannot be empty")
	}
	if id <= 0 {
		return errors.New("actor id should be greater than zero")
	}
	return as.store.UpdateActor(id, actor)
}

// Удаляет актера по ID

func (as *ActorService) DeleteActor(id int) error {
	if id <= 0 {
		return errors.New("actor id should be greater than zero")
	}
	return as.store.DeleteActor(id)
}
