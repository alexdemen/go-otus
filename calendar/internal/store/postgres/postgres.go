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
			returning id
		`

	result, err := p.database.QueryContext(cxt, sql, event.Name, *event.Description, event.StartDate, event.Duration)
	if err != nil {
		return event, err
	}
	defer result.Close()

	if result.Next() {
		err = result.Scan(&event.Id)
	} else {
		return event, errors.New("no get last inserted id")
	}

	return event, nil
}

func (p Store) Edit(cxt context.Context, event core.Event) error {
	sql := ``
	result, err := p.database.ExecContext(cxt, sql, event.Name, event.Description, event.StartDate, event.Duration, event.Id)
	if err != nil {
		return err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return err
	} else if count == 0 {
		return errors.New("now affected row")
	}

	return nil
}

func (p Store) Remove(cxt context.Context, event core.Event) error {
	sql := ``

	_, err := p.database.ExecContext(cxt, sql, event.Id)
	return err
}

func (p Store) List(cxt context.Context) ([]core.Event, error) {
	sql := `select id, name, description, date, duration
			from events
			where deleted = false`

	result, err := p.database.QueryContext(cxt, sql)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	events := make([]core.Event, 0)

	for result.Next() {
		event := core.Event{}
		err = result.Scan(&event.Id, &event.Name, event.Description, event.StartDate, event.Duration)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}