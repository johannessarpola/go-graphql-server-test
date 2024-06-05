package app

import (
	"github.com/johannessarpola/go-graphql-server-test/pkg/spotify"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	SpotifyConfig spotify.Config `yaml:"spotify"`
	Port          string         `yaml:"port"`
}

func LoadConfig[T interface{}](filename string) (T, error) {
	var config T

	configFile, err := os.ReadFile(filename)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
