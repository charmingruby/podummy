package pokeapi

import "github.com/charmingruby/podummy/internal/poke"

type Ability struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type AbilityDetail struct {
	Ability  Ability `json:"ability"`
	IsHidden bool    `json:"is_hidden"`
	Slot     int     `json:"slot"`
}

type Cry struct {
	Latest string `json:"latest"`
	Legacy string `json:"legacy"`
}

type Form struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Version struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type GameIndex struct {
	GameIndex int     `json:"game_index"`
	Version   Version `json:"version"`
}

type Move struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type VersionGroupDetail struct {
	LevelLearnedAt  int     `json:"level_learned_at"`
	MoveLearnMethod Method  `json:"move_learn_method"`
	VersionGroup    Version `json:"version_group"`
}

type Method struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type MoveDetail struct {
	Move                Move                 `json:"move"`
	VersionGroupDetails []VersionGroupDetail `json:"version_group_details"`
}

type Sprite struct {
	BackDefault  string `json:"back_default"`
	BackShiny    string `json:"back_shiny"`
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}

type OtherSprites struct {
	DreamWorld struct {
		FrontDefault string `json:"front_default"`
	} `json:"dream_world"`
	Home struct {
		FrontDefault string `json:"front_default"`
	} `json:"home"`
}

type Species struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Stat struct {
	BaseStat int `json:"base_stat"`
	Effort   int `json:"effort"`
	Stat     struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"stat"`
}

type Type struct {
	Slot int `json:"slot"`
	Type struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"type"`
}

type Pokemon struct {
	Abilities              []AbilityDetail `json:"abilities"`
	BaseExperience         int             `json:"base_experience"`
	Cries                  Cry             `json:"cries"`
	Forms                  []Form          `json:"forms"`
	GameIndices            []GameIndex     `json:"game_indices"`
	Height                 int             `json:"height"`
	HeldItems              []interface{}   `json:"held_items"`
	ID                     int             `json:"id"`
	IsDefault              bool            `json:"is_default"`
	LocationAreaEncounters string          `json:"location_area_encounters"`
	Moves                  []MoveDetail    `json:"moves"`
	Sprites                Sprite          `json:"sprites"`
	Species                Species         `json:"species"`
	Stats                  []Stat          `json:"stats"`
	Types                  []Type          `json:"types"`
	Weight                 int             `json:"weight"`
	Name                   string          `json:"name"`
	Order                  int             `json:"order"`
	PastAbilities          []interface{}   `json:"past_abilities"`
	PastTypes              []interface{}   `json:"past_types"`
}

func (p *Pokemon) ToDomain() poke.Pokemon {
	return poke.Pokemon{
		ID:     p.ID,
		Name:   p.Name,
		Sprite: p.Sprites.FrontDefault,
	}
}
