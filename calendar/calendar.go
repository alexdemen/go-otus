package main

import (
	"errors"
	"fmt"
	"github.com/alexdemen/go-otus/calendar/internal/config"
	"github.com/alexdemen/go-otus/calendar/internal/middleware"
	flag "github.com/spf13/pflag"
	"log"
	"net/http"
)

var configPath = flag.String("config", "", "path to configuration file")

func init() {
	flag.Parse()
}

func main() {
	runConfig, err := config.GetConfiguration(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	level, err := chooseLoggerLevel(runConfig.LogLevel)
	if err != nil {
		log.Fatal(err)
	}

	err = middleware.ConfigureLogger(runConfig.LogFile, level)
	if err != nil {
		log.Fatal(err)
	}

	handler := middleware.SetLogger(http.HandlerFunc(resolvePath))
	if err := http.ListenAndServe(runConfig.ListenAddress, handler); err != nil {
		fmt.Println(err)
	}

	middleware.CloseLogger()
}

func resolvePath(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/hello":
		w.Write([]byte("Hello world!"))
		return
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func chooseLoggerLevel(levelName string) (int, error) {
	switch levelName {
	case "debug":
		return middleware.Debug, nil
	case "warn":
		return middleware.Warning, nil
	case "error":
		return middleware.Error, nil
	case "info":
		return middleware.Info, nil
	default:
		return 0, errors.New("failed to select logger level")
	}
}
