package repository

import (
	"context"
	_ "embed"

	"github.com/nleof/goyesql"
	"github.com/rafaelrubbioli/espenses/pkg/entity"
	"github.com/rafaelrubbioli/espenses/pkg/lib/db"
)

type User interface {
	Create(ctx context.Context, name, password string) (int64, error)
	Get(ctx context.Context, id int64) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id int64) error
}

//go:embed user.sql
var userQueries []byte

type user struct {
	db      db.MySQL
	queries goyesql.Queries
}

func NewUserRepository(db db.MySQL) User {
	return user{
		db:      db,
		queries: goyesql.MustParseBytes(userQueries),
	}
}

func (r user) Create(ctx context.Context, name, password string) (int64, error) {
	result, err := r.db.Exec(ctx, r.queries["create"], name, password)
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
	err := r.db.Get(ctx, &user, r.queries["by-id"], id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r user) Update(ctx context.Context, user *entity.User) error {
	_, err := r.db.Exec(ctx, r.queries["update"], user.Name, user.ID)
	return err
}

func (r user) Delete(ctx context.Context, id int64) error {
	_, err := r.db.Exec(ctx, r.queries["delete"], id)
	return err
}
