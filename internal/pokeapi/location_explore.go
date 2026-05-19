package pokeapi

func (c *Client) GetAreaPokemon(areaName string) (AreaPokemonResponse, error) {
	url := baseURL + "/location-area/" + areaName
	return getAndCache[AreaPokemonResponse](c, url)
}
