package controller

import "time"

type serviceActor interface {
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
