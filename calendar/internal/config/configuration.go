package config

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Configuration struct {
	ListenAddress string `yaml:"http_listen"`
	LogFile       string `yaml:"log_file"`
	LogLevel      string `yaml:"log_level"`
	DSN           string `yaml:"database_url"`
}

func GetConfiguration(configPath string) (Configuration, error) {
	configuration := Configuration{}
	err := checkConfigFile(configPath)
	if err != nil {
		return configuration, err
	}

	configData, err := ioutil.ReadFile(configPath)
	if err != nil {
		return configuration, err
	}

	err = yaml.Unmarshal(configData, &configuration)
	if err != nil {
		return configuration, err
	}

	return configuration, nil
}

func checkConfigFile(configPath string) error {
	if configPath == "" {
		return errors.New("no configuration file specified")
	}

	if _, err := os.Stat(configPath); err != nil {
		return err
	}

	return nil
}
