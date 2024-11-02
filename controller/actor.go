package controller

import "time"

type ServiceActor interface {
	CreateActor(name string, birthday time.Time, gender string) (int, error)
	GetActor(id int) (Actor, error)
	UpdateActor(id int, name string, birthday time.Time, gender string) (int, error)
	DeleteActor(id int) error
}
type Actor struct {
	ID       int
	Name     string
	Birthday time.Time
	Gender   string
}
type actorController struct {
	service ServiceActor
}

func NewActorController(service ServiceActor) *actorController {
	return &actorController{service: service}
}

func (ac *actorController) CreateActor(name string, birthday time.Time, gender string) (int, error) {
	return ac.service.CreateActor(name, birthday, gender)
}

func (ac *actorController) GetActor(id int) (Actor, error) {
	return ac.service.GetActor(id)
}

func (ac *actorController) UpdateActor(id int, name string, birthday time.Time, gender string) (int, error) {
	return ac.service.UpdateActor(id, name, birthday, gender)
}

func (ac *actorController) Delete(id int) error {
	return ac.service.DeleteActor(id)
}
