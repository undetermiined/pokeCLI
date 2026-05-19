package pokeapi

type NamedAPI struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationAreasResponse struct {
	Next     *string    `json:"next"`
	Previous *string    `json:"previous"`
	Results  []NamedAPI `json:"results"`
}
