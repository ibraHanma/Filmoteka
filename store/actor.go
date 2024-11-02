package store

import (
	"Filmoteka/service"
	"database/sql"
	"fmt"
)

type actor struct {
	db *sql.DB
}

func NewActor(db *sql.DB) *actor {
	return &actor{db: db}
}

type ActorStore struct {
	actor []service.Actor
}

func (r *ActorStore) CreateActor(actor service.Actor) error {
	r.actor = append(r.actor, actor)

	return nil

}
func (r *ActorStore) GetActor(id int) (*service.Actor, error) {
	for _, actor := range r.actor {
		if actor.ID == id {
			return &actor, nil
		}
	}
	return nil, fmt.Errorf("актер не найден")

}
func (r *ActorStore) UpdateActor(actor *service.Actor) error {
	for i, m := range r.actor {
		if m.ID == actor.ID {
			r.actor[i] = *actor
			return nil
		}
	}
	return fmt.Errorf("актер не найден")

}
func (r *ActorStore) DeleteActor(id int) error {

	for i, m := range r.actor {
		if m.ID == id {
			r.actor = append(r.actor[:i], r.actor[i+1:]...)
		}
	}
	return fmt.Errorf("актер не найден")
}
