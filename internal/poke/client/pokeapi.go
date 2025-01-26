package client

import "github.com/charmingruby/podummy/internal/poke"

type PokeAPI interface {
	GetPokemonByID(id string) (poke.Pokemon, error)
}
