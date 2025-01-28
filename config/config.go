package config

import (
	env "github.com/caarlos0/env/v6"
)

type environment struct {
	Version string `env:"VERSION,required"`
}

func New() (*Wrapper, error) {
	environment := environment{}
	if err := env.Parse(&environment); err != nil {
		return &Wrapper{}, err
	}

	return &Wrapper{
		Versioning: versioning(environment),
	}, nil
}

type Wrapper struct {
	Versioning versioning
}

type versioning struct {
	Version string
}
