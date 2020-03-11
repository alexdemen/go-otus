package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/alexdemen/go-otus/calendar/internal/core"
	_ "github.com/jackc/pgx/stdlib"
)

type Store struct {
	database *sql.DB
}

func NewStore(dsn string) (*Store, error) {
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

	return &Store{database: db}, nil
}

func (p Store) Add(cxt context.Context, event core.Event) (core.Event, error) {
	sql := `insert into events (name, description, date, duration) 
			values ($1, $2, $3, $4)
			returning id`

	result, err := p.database.QueryContext(cxt, sql, event.Name, event.Description, event.StartDate, event.Duration)
	if err != nil {
		return event, err
	}

	if result.Next() {
		err = result.Scan(&event.Id)
	} else {
		return event, errors.New("no get last inserted id")
	}

	return event, nil
}

func (p Store) Edit(cxt context.Context, event core.Event) error {
	panic("implement me")
}

func (p Store) Remove(cxt context.Context, event core.Event) error {
	panic("implement me")
}

func (p Store) List(cxt context.Context) ([]core.Event, error) {
	panic("implement me")
}
