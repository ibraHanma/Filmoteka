package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"time"
)

type Actor struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Birthday time.Time `json:"birthday"`
	Gender   string    `json:"gender"`
}
type ActorRepo struct {
	db *sql.DB
}

func NewActor(db *sql.DB) *ActorRepo {
	return &ActorRepo{db: db}
}

func (a *ActorRepo) CreateActor(actor Actor) (int, error) {
	var id int

	query, args, err := squirrel.Insert("actor").
		Columns("name", "birthday", "gender").
		Values(actor.Name, actor.Birthday, actor.Gender).
		Suffix("RETURNING id").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return 0, fmt.Errorf("не удалось создать SQL запрос: %w", err)
	}

	err = a.db.QueryRow(query, args...).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("не удалось выполнить запрос: %w", err)
	}
	return id, nil
}

func (a *ActorRepo) GetActor(id int) (Actor, error) {
	var actor Actor

	query, args, err := squirrel.Select("id", "name", "birthday", "gender").
		From("actor").
		Where(squirrel.Eq{"id": id}).
		ToSql()

	if err != nil {
		return Actor{}, fmt.Errorf("не удалось создать SQL запрос: %w", err)
	}

	err = a.db.QueryRow(query, args...).Scan(&actor.ID, &actor.Name, &actor.Birthday, &actor.Gender)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Actor{}, fmt.Errorf("актер с идентификатором %d не найден", id)
		}
		return Actor{}, fmt.Errorf("не удалось выполнить запрос: %w", err)
	}
	return actor, nil
}

func (a *ActorRepo) UpdateActor(actor Actor) error {
	_, err := a.GetActor(actor.ID)
	if err != nil {
		return err
	}

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

func (a *ActorRepo) DeleteActor(id int) error {
	_, err := a.GetActor(id)
	if err != nil {
		return err
	}

	query, args, err := squirrel.Delete("actor").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return fmt.Errorf("не удалось создать SQL запрос: %w", err)
	}

	_, err = a.db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("не удалось удалить актера с идентификатором %d: %w", id, err)
	}
	return nil
}
