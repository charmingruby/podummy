package poke

import "github.com/charmingruby/podummy/pkg/helper"

const (
	TEAM_SIZE = 6
)

func NewService(version string, pokeAPIClient PokeAPIClient) *Service {
	return &Service{
		Version:       version,
		PokeAPIClient: pokeAPIClient,
	}
}

type Service struct {
	Version       string
	PokeAPIClient PokeAPIClient
}

type TeamShuffledResponse struct {
	CurrentVersion string    `json:"current_version"`
	Team           []Pokemon `json:"team_shuffled"`
}

func (s *Service) ShuffleTeam() (TeamShuffledResponse, error) {
	opts, err := retrieveOptsFromVersion(s.Version)
	if err != nil {
		return TeamShuffledResponse{}, err
	}

	team := make([]Pokemon, 0, TEAM_SIZE)
	usedIDs := make(map[string]bool)

	for len(team) < TEAM_SIZE {
		id := helper.GenerateRandomValueFromRange(opts.MaxID, opts.MinID)

		if usedIDs[id] {
			continue
		}

		pokemon, err := s.PokeAPIClient.GetPokemonByID(id)
		if err != nil {
			return TeamShuffledResponse{}, err
		}

		team = append(team, pokemon)
		usedIDs[id] = true
	}

	return TeamShuffledResponse{
		CurrentVersion: s.Version,
		Team:           team,
	}, nil
}
