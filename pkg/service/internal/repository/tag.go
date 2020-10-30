package repository

import (
	"context"

	"github.com/rafaelrubbioli/espenses/pkg/entity"
	"github.com/rafaelrubbioli/espenses/pkg/lib/db"
)

type Tag interface{
	Create(ctx context.Context, name string) (int64, error)
	Get(ctx context.Context, id int64) (*entity.Tag, error)
	Update(ctx context.Context, user *entity.Tag) error
	Delete(ctx context.Context, id int64) error
}

type tag struct {
	db db.MySQL
}

func NewTagRepository(db db.MySQL) Tag {
	return tag{db: db}
}

func (r tag) Create(ctx context.Context, name string) (int64, error) {
	query := `INSERT INTO tags (name) VALUES (?)`
	result, err := r.db.Exec(ctx, query, name)
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
	query := `SELECT id, name FROM tags WHERE id = ? LIMIT 1`
	err := r.db.Get(ctx, &tag, query, id)
	if err != nil {
		return nil, err
	}

	return &tag, nil
}

func (r tag) Update(ctx context.Context, tag *entity.Tag) error {
	query := `UPDATE tags SET name = (?) WHERE id = ?`
	_, err := r.db.Exec(ctx, query, tag.Name, tag.ID)
	return err
}

func (r tag) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM tags WHERE id = ?`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
