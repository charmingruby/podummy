package poke

type PokeAPIClient interface {
	GetPokemonByID(id string) (Pokemon, error)
}
