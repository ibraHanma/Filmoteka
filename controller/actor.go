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
type ActorController struct {
	service ServiceActor
}

func NewActorController(service ServiceActor) *ActorController {
	return &ActorController{service: service}
}

func (ac *ActorController) CreateActor(name string, birthday time.Time, gender string) (int, error) {
	return ac.service.CreateActor(name, birthday, gender)
}

func (ac *ActorController) GetActor(id int) (Actor, error) {
	return ac.service.GetActor(id)
}

func (ac *ActorController) UpdateActor(id int, name string, birthday time.Time, gender string) (int, error) {
	return ac.service.UpdateActor(id, name, birthday, gender)
}

func (ac *ActorController) Delete(id int) error {
	return ac.service.DeleteActor(id)
}
