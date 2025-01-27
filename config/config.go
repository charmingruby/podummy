package config

import (
	env "github.com/caarlos0/env/v6"
)

type environment struct {
	ServerPort string `env:"SERVER_PORT,required"`
	Version    string `env:"VERSION,required"`
}

func New() (*Wrapper, error) {
	environment := environment{}
	if err := env.Parse(&environment); err != nil {
		return &Wrapper{}, err
	}

	return &Wrapper{
		Versioning: versioning{
			Version: environment.Version,
		},
		Server: server{
			Port: environment.ServerPort,
		},
	}, nil
}

type Wrapper struct {
	Versioning versioning
	Server     server
}

type versioning struct {
	Version string
}

type server struct {
	Port string
}
