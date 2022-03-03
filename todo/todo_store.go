package todo

import (
	"context"
	"errors"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type TodoStorePG struct {
	db *pgxpool.Pool
}

func NewTodoStorePG(db *pgxpool.Pool) *TodoStorePG {
	return &TodoStorePG{db}
}

func (store TodoStorePG) Add(todo Todo) (*Todo, error) {
	_, err := store.db.Exec(
		context.Background(),
		"INSERT INTO \"Todo\" (id, text, user_id) VALUES ($1, $2, $3)",
		todo.ID,
		todo.Text,
		todo.UserId,
	)

	if err != nil { return nil, errors.New(err.Error()) }
	return &todo, nil
}

func (store TodoStorePG) Delete(todoID uuid.UUID) error {
	return nil
}

func (store TodoStorePG) Toggle(todoID uuid.UUID, done bool) (*Todo, error) {
	return nil, nil
}

func (store TodoStorePG) UpdateText(todoID uuid.UUID, text string) (*Todo, error) {
	return nil, nil
}

func (store TodoStorePG) FindByID(todoID uuid.UUID) (*Todo, error) {
	return nil, nil
}

func (store TodoStorePG) FindByUserID(userID uuid.UUID) ([]*Todo, error) {
	return nil, nil
}