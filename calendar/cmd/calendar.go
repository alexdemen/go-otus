package main

import (
	"errors"
	"fmt"
	"github.com/alexdemen/go-otus/calendar/internal/config"
	flag "github.com/spf13/pflag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var configPath = flag.String("config", "", "path to configuration file")

func init() {
	flag.Parse()
}
func main() {
	runConfig, err := getConfiguration()
	if err != nil {
		log.Fatal(err)
	}

	handler := http.HandlerFunc(final)
	newHandler := logMiddleWare(handler)
	http.Handle("/", newHandler)
	if err := http.ListenAndServe(runConfig.ListenAddress, nil); err != nil {
		fmt.Println(err)
	}
}

func logMiddleWare(next http.Handler) http.Handler {
	return next
}

func final(w http.ResponseWriter, r *http.Request) {
}

func getConfiguration() (*config.Configuration, error) {
	err := checkConfigFile()
	if err != nil {
		return nil, err
	}

	configData, err := ioutil.ReadFile(*configPath)
	if err != nil {
		return nil, err
	}

	configuration := config.Configuration{}
	err = yaml.Unmarshal(configData, &configuration)
	if err != nil {
		return nil, err
	}

	return &configuration, nil
}

func checkConfigFile() error {
	if *configPath == "" {
		return errors.New("no configuration file specified")
	}

	if _, err := os.Stat(*configPath); err != nil {
		return err
	}

	return nil
}
