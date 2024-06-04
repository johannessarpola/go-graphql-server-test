package app

import (
	"github.com/johannessarpola/graphql-test/pkg/spotify"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	SpotifyConfig spotify.Config `yaml:"spotify"`
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
