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

type QueryConfig struct {
	QueryUrl string `yaml:"rabbit_url"`
}

func GetConfiguration(config interface{}, configPath string) error {
	err := checkConfigFile(configPath)
	if err != nil {
		return err
	}

	configData, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		return err
	}

	return nil
}

func GetQueryConfiguration(config *QueryConfig, configPath string) error {
	err := checkConfigFile(configPath)
	if err != nil {
		return err
	}

	configData, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(configData, config)
	if err != nil {
		return err
	}

	return nil
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
