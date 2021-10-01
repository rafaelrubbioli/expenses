package repository

import (
	"context"
	_ "embed"

	"github.com/nleof/goyesql"

	"github.com/rafaelrubbioli/espenses/pkg/entity"
	"github.com/rafaelrubbioli/espenses/pkg/lib/db"
)

type Tag interface {
	Create(ctx context.Context, name string) (int64, error)
	Get(ctx context.Context, id int64) (*entity.Tag, error)
	Update(ctx context.Context, user *entity.Tag) error
	Delete(ctx context.Context, id int64) error
}

//go:embed tag.sql
var tagQueries []byte

type tag struct {
	db      db.MySQL
	queries goyesql.Queries
}

func NewTagRepository(db db.MySQL) Tag {
	return tag{
		db:      db,
		queries: goyesql.MustParseBytes(tagQueries),
	}
}

func (r tag) Create(ctx context.Context, name string) (int64, error) {
	result, err := r.db.Exec(ctx, r.queries["create"], name)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return id, nil
}

func (r tag) Get(ctx context.Context, id int64) (*entity.Tag, error) {
	var tag entity.Tag
	err := r.db.Get(ctx, &tag, r.queries["by-id"], id)
	if err != nil {
		return nil, err
	}

	return &tag, nil
}

func (r tag) Update(ctx context.Context, tag *entity.Tag) error {
	_, err := r.db.Exec(ctx, r.queries["update"], tag.Name, tag.ID)
	return err
}

func (r tag) Delete(ctx context.Context, id int64) error {
	_, err := r.db.Exec(ctx, r.queries["delete"], id)
	return err
}
