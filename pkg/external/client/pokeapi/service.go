package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/charmingruby/podummy/internal/poke"
)

const (
	BASE_URL = "https://pokeapi.co/api/v2"
)

type PokeAPI struct{}

func NewPokeAPI() *PokeAPI {
	return &PokeAPI{}
}

func (a *PokeAPI) GetPokemonByID(id string) (poke.Pokemon, error) {
	url := BASE_URL + "/pokemon/" + id

	res, err := http.Get(url)
	if err != nil {
		return poke.Pokemon{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return poke.Pokemon{}, err
	}

	var data *Pokemon
	if err := json.Unmarshal(body, &data); err != nil {
		return poke.Pokemon{}, err
	}

	return data.ToDomain(), nil
}
