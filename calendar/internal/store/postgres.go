package store

import (
	"context"
	"database/sql"
	"github.com/alexdemen/go-otus/calendar/internal/core"
	_ "github.com/jackc/pgx/stdlib"
)

type Postgres struct {
	database *sql.DB
}

func NewPostgres(dsn string) (*Postgres, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(5)

	return &Postgres{database: db}, nil
}

func (p Postgres) Add(cxt context.Context, event core.Event) (core.Event, error) {
	panic("implement me")
}

func (p Postgres) Edit(cxt context.Context, event core.Event) error {
	panic("implement me")
}

func (p Postgres) Remove(cxt context.Context, event core.Event) error {
	panic("implement me")
}

func (p Postgres) List(cxt context.Context) ([]core.Event, error) {
	panic("implement me")
}
