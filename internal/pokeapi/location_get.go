package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetAreaPokemon(areaName string) (AreaPokemonResponse, error) {
	url := baseURL + "/location-area/" + areaName

	if val, ok := c.cache.Get(url); ok {
		pokemonResponse := AreaPokemonResponse{}
		err := json.Unmarshal(val, &pokemonResponse)
		if err != nil {
			return AreaPokemonResponse{}, err
		}
		return pokemonResponse, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return AreaPokemonResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return AreaPokemonResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return AreaPokemonResponse{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return AreaPokemonResponse{}, err
	}

	pokemonResponse := AreaPokemonResponse{}
	err = json.Unmarshal(dat, &pokemonResponse)
	if err != nil {
		return AreaPokemonResponse{}, err
	}

	c.cache.Add(url, dat)
	return pokemonResponse, nil
}
