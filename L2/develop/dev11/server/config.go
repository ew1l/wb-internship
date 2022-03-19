package server

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

// Config for server
type Config struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

// NewConfig config's structure constructor
func NewConfig() *Config {
	config := new(Config)

	file, _ := ioutil.ReadFile("config/config.yml")
	if err := yaml.Unmarshal(file, config); err != nil {
		log.Println(err)
	}

	return &Config{
		Host: config.Host,
		Port: config.Port,
	}
}
