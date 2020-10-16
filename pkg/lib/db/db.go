package db

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type MySQL interface {
	Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

func NewMySQL(db *sqlx.DB) MySQL {
	return mysql{
		db: db,
	}
}

type mysql struct {
	db *sqlx.DB
}

func (m mysql) Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return m.db.GetContext(ctx, dest, query, args)
}
func (m mysql) Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return m.db.SelectContext(ctx, dest, query, args)
}

func (m mysql) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return m.db.ExecContext(ctx, query, args)
}
