package main

import (
	"context"
	"github.com/alexdemen/go-otus/calendar/internal/config"
	"github.com/alexdemen/go-otus/calendar/internal/scheduler"
	"github.com/alexdemen/go-otus/calendar/internal/store/postgres"
	"github.com/spf13/pflag"
	"log"
	"os"
	"os/signal"
	"time"
)

var configPath = pflag.String("config", "", "path to configuration file")

func init() {
	pflag.Parse()
}

func main() {
	runConfig := config.SchedulerConfig{}
	err := config.GetSchedulerConfiguration(&runConfig, *configPath)
	if err != nil {
		log.Fatal(err)
	}

	explorer, err := postgres.NewStore(runConfig.DSN)
	if err != nil {
		log.Fatal(err)
	}

	sched := scheduler.NewScheduler(explorer, runConfig.QueryUrl)

	ctx, cancel := context.WithCancel(context.Background())

	done := make(chan os.Signal)
	signal.Notify(done, os.Interrupt)

	go sched.Run(ctx)

	<-done
	cancel()
	time.Sleep(5 * time.Second)
}
