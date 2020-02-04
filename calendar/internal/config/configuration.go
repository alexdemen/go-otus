package config

type Configuration struct {
	ListenAddress string `yaml:"http_listen"`
	LogFile       string `yaml:"log_file"`
	LogLevel      string `yaml:"log_level"`
}
