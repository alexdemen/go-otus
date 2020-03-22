package main

import (
	"fmt"
	"github.com/alexdemen/go-otus/calendar/internal/config"
	"github.com/spf13/pflag"
	"github.com/streadway/amqp"
	"log"
	"os"
	"os/signal"
)

var configPath = pflag.String("config", "", "path to configuration file")

func init() {
	pflag.Parse()
}

func main() {
	configuration := config.QueryConfig{}
	err := config.GetQueryConfiguration(&configuration, *configPath)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := amqp.Dial(configuration.QueryUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	channel, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}

	done := make(chan os.Signal)
	signal.Notify(done, os.Kill, os.Interrupt)

	msgs, err := channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil)

	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for msg := range msgs {
			fmt.Println(string(msg.Body))
		}
	}()

	<-done
}
