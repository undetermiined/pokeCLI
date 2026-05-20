package pokeapi

// location_explore.go
type PokemonEncounter struct {
	Pokemon NamedAPI `json:"pokemon"`
}

type AreaPokemonResponse struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}
