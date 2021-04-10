package repository

import (
	"context"

	"github.com/rafaelrubbioli/espenses/pkg/entity"
	"github.com/rafaelrubbioli/espenses/pkg/lib/db"
)

type User interface{
	Create(ctx context.Context, name , password string) (int64, error)
	Get(ctx context.Context, id int64) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id int64) error
}

type user struct {
	db db.MySQL
}

func NewUserRepository(db db.MySQL) User {
	return user{db: db}
}

func (r user) Create(ctx context.Context, name, password string) (int64, error) {
	query := `INSERT INTO users (name, password) VALUES (?, ?)`
	result, err := r.db.Exec(ctx, query, name, password)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return id, nil
}

func (r user) Get(ctx context.Context, id int64) (*entity.User, error) {
	var user entity.User
	query := `SELECT id, name FROM users WHERE id = ? LIMIT 1`
	err := r.db.Get(ctx, &user, query, id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r user) Update(ctx context.Context, user *entity.User) error {
	query := `UPDATE users SET name = (?) WHERE id = ?`
	_, err := r.db.Exec(ctx, query, user.Name, user.ID)
	return err
}

func (r user) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM users WHERE id = ?`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
