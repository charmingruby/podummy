package poke

import (
	"math/rand"
	"strconv"
	"time"
)

const (
	TEAM_SIZE = 6

	MIN_ID = 1
	MAX_ID = 151
)

func NewService(pokeAPIClient PokeAPIClient) *Service {
	return &Service{
		PokeAPIClient: pokeAPIClient,
	}
}

type Service struct {
	PokeAPIClient PokeAPIClient
}

func (s *Service) ShuffleTeam() ([]Pokemon, error) {
	team := make([]Pokemon, 0, TEAM_SIZE)
	usedIDs := make(map[string]bool)

	for len(team) < TEAM_SIZE {
		id := s.generateRandomValidID()

		if usedIDs[id] {
			continue
		}

		pokemon, err := s.PokeAPIClient.GetPokemonByID(id)
		if err != nil {
			return nil, err
		}

		team = append(team, pokemon)
		usedIDs[id] = true
	}

	return team, nil
}

func (s *Service) generateRandomValidID() string {
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)

	id := rng.Intn(MAX_ID-MIN_ID) + MIN_ID

	return strconv.Itoa(id)
}
