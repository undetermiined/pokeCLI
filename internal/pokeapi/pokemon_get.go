package pokeapi

func (c *Client) GetPokemon(pokeName string) (PokemonResponse, error) {
	url := baseURL + "/pokemon/" + pokeName
	return getAndCache[PokemonResponse](c, url)
}
