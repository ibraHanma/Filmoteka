package repository

import (
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
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

	query, args, err := squirrel.Insert("actor").
		Columns("name", "birthday", "gender").
		Values(actor.Name, actor.Birthday, actor.Gender).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return 0, err
	}
	err = a.db.QueryRow(query, args...).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("не удалось выполнить запрос:%w", err)
	}
	return id, nil
}

func (a *actor) GetActor(id int) (Actor, error) {
	var actor Actor

	query, args, err := squirrel.Select("id", "name", "birthday", "gender").
		From("actor").
		Where(squirrel.Eq{"id": id}).
		ToSql()

	if err != nil {
		return Actor{}, err
	}

	err = a.db.QueryRow(query, args...).Scan(&actor.ID, &actor.Name, &actor.Birthday, &actor.Gender)
	if err != nil {
		if err == sql.ErrNoRows {
			return Actor{}, fmt.Errorf("актер с идентификатором не найден", id)
		}
		return Actor{}, err
	}
	return actor, nil
}

func (a *actor) UpdateActor(actor Actor) error {
	query, args, err := squirrel.Update("actor").
		Set("name", actor.Name).
		Set("birthday", actor.Birthday).
		Set("gender", actor.Gender).
		Where(squirrel.Eq{"id": actor.ID}).
		ToSql()
	if err != nil {
		return fmt.Errorf("не удалось создать SQL запрос: %w", err)
	}
	_, err = a.db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("не удалось обновить актера с ID %d: %w", actor.ID, err)

	}
	return nil

}

func (a *actor) DeleteActor(id int) error {
	query, args, err := squirrel.Delete("actor").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return fmt.Errorf("не удалось создать SQL запрос: %w", err)
	}
	_, err = a.db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("не удалось удалить участника с идентификатором %d: %w", id, err)
	}
	return nil

}
