package scheduler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/alexdemen/go-otus/calendar/internal/core"
	"github.com/streadway/amqp"
	"time"
)

type Scheduler struct {
	source core.Explorer
	url    string
}

func NewScheduler(source core.Explorer, url string) *Scheduler {
	return &Scheduler{source: source, url: url}
}

func (s *Scheduler) Run(cxt context.Context) {
	ticker := time.Tick(30 * time.Second)

	go func() {
		for {
			select {
			case <-ticker:
				err := notify(cxt, s.source, s.url)
				if err != nil {
					fmt.Println(err)
				}
			case <-cxt.Done():
				return
			}
		}
	}()
}

func notify(ctx context.Context, source core.Explorer, url string) error {
	data, err := source.Get(ctx)
	if err != nil {
		return err
	}

	conn, err := amqp.Dial(url)
	if err != nil {
		return err
	}
	defer conn.Close()

	channel, err := conn.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	q, err := channel.QueueDeclare(
		"event",
		false,
		false,
		false,
		false,
		nil)
	if err != nil {
		return err
	}

	for _, msg := range data {
		body, err := json.Marshal(msg)
		if err != nil {
			return err
		}

		err = channel.Publish("", q.Name, false, false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        body,
			})

		if err != nil {
			return err
		}
	}

	return nil
}
