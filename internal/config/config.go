package config

import (
	"gitlab.com/bongerka/lg"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	HTTPServer *HttpServerConfig `yaml:"HTTP_SERVER"`
	Repo       *RepoConfig       `yaml:"REPOSITORY"`
}

type RepoConfig struct {
	Size int `yaml:"SIZE_MAP"`
}

type HttpServerConfig struct {
	Addr string `yaml:"ADDR"`
}

func MustParse(cfg *Config) {
	file, err := os.Open("./config.yml")
	if err != nil {
		lg.Fatalf("unable to open config file: %v", err)
	}

	if err = yaml.NewDecoder(file).Decode(cfg); err != nil {
		lg.Fatalf("unable to decode config file: %v", err)
	}
}
