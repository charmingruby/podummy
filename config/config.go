package config

import "github.com/charmingruby/podummy/internal/poke"

func New() (*Wrapper, error) {
	return &Wrapper{
		Versioning: versioning{
			Version: poke.GREEN_VERSION,
		},
		Server: server{
			Port: "8080",
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
