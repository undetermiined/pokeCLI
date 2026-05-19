package pokeapi

// location_explore.go
type PokemonEncounter struct {
	Pokemon NamedAPI `json:"pokemon"`
}

type AreaPokemonResponse struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

// pokemon_get.go
type PokemonResponse struct {
	ID                     int        `json:"id"`
	IsDefault              bool       `json:"is_default"`
	LocationAreaEncounters string     `json:"location_area_encounters"`
	Name                   string     `json:"name"`
	Order                  int        `json:"order"`
	BaseExperience         int        `json:"base_experience"`
	Forms                  []NamedAPI `json:"forms"`
	Types                  []struct {
		Slot int      `json:"slot"`
		Type NamedAPI `json:"type"`
	} `json:"types"`
}
