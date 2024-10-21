package repository

import (
	"database/sql"
	"time"
)

type Actor struct {
	ID       int
	Name     string
	Birthday time.Time
	Gender   string
}

type actor struct {
	db *sql.DB
}

func NewActor(db *sql.DB) *actor {
	return &actor{db: db}

}

func (a *actor) CreateActor(actor Actor) (int, error) {
	var id int
	query := `INSERT INTO actor (name, birthday, gender) VALUES ($1, $2, $3) RETURNING id`
	err := a.db.QueryRow(query, actor.Name, actor.Birthday, actor.Gender).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (a *actor) GetActorByID(id int) (Actor, error) {
	var actor Actor
	query := `SELECT id,name,birthday,gender FROM actor WHERE id = $1`
	err := a.db.QueryRow(query, id).Scan(&actor.ID, &actor.Name, &actor.Birthday, &actor.Gender)
	if err != nil {
		return Actor{}, err
	}
	return actor, err

}

func (a *actor) UpdateActor(actor Actor) error {
	query := `UPDATE actor SET name = $1, birthday = $2, gender = $3 WHERE id = $4`
	_, err := a.db.Exec(query, actor.Name, actor.Birthday, actor.Gender, actor.ID)
	return err
}

func (a *actor) DeleteActor(id int) error {
	query := `DELETE FROM actor WHERE id = $1`
	_, err := a.db.Exec(query, id)
	return err
}
