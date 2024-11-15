package store

import (
	"Filmoteka/internal/repository"
	"Filmoteka/internal/service"
)

type Actor struct {
	repo *repository.ActorRepo // или как вы назвали свой репозиторий
}

func (a Actor) CreateActor(actor service.Actor) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (a Actor) GetActor(id int) (service.Actor, error) {
	//TODO implement me
	panic("implement me")
}

func (a Actor) UpdateActor(id int, actor service.Actor) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) DeleteActor(id int) error {
	//TODO implement me
	panic("implement me")
}

func NewActor(repo *repository.ActorRepo) *Actor {
	return &Actor{repo: repo}
}

type ActorStore struct {
	actor []service.Actor
}

//
//func (r *ActorStore) CreateActor(actor service.Actor) error {
//	r.actor = append(r.actor, actor)
//
//	return nil
//}
//
//func (r *ActorStore) GetActor(id int) (*service.Actor, error) {
//	for _, actor := range r.actor {
//		if actor.ID == id {
//			return &actor, nil
//		}
//	}
//	return nil, fmt.Errorf("actor not found")
//}
//
//func (r *ActorStore) UpdateActor(actor *service.Actor) error {
//	for i, m := range r.actor {
//		if m.ID == actor.ID {
//			r.actor[i] = *actor
//
//			return nil
//		}
//	}
//	return fmt.Errorf("actor not found")
//}
//
//func (r *ActorStore) DeleteActor(id int) error {
//	for i, m := range r.actor {
//		if m.ID == id {
//			r.actor = append(r.actor[:i], r.actor[i+1:]...)
//
//			return nil
//		}
//	}
//	return fmt.Errorf("actor not found")
//}
