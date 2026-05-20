package pokeapi

// pokemon_get.go
type PokemonResponse struct {
	Name                   string  `json:"name"`
	Order                  int     `json:"order"`
	BaseExperience         int     `json:"base_experience"`
	Height                 int     `json:"height"`
	Weight                 int     `json:"weight"`
	ID                     int     `json:"id"`
	LocationAreaEncounters string  `json:"location_area_encounters"`
	Stats                  []Stats `json:"stats"`
	Types                  []Types `json:"types"`
	Forms                  []Forms `json:"forms"`
}
type Forms struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Stat struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Stats struct {
	BaseStat int  `json:"base_stat"`
	Effort   int  `json:"effort"`
	Stat     Stat `json:"stat"`
}
type Type struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Types struct {
	Slot int  `json:"slot"`
	Type Type `json:"type"`
}
