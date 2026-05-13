package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationAreasResponse struct {
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationArea `json:"results"`
}

func GetLocationAreas(pageUrl string) (LocationAreasResponse, error) {
	var result LocationAreasResponse
	resp, err := http.Get("https://pokeapi.co/api/v2/location-area")
	if err != nil {
		return LocationAreasResponse{}, fmt.Errorf("location-area request failed: %w", err)
	}
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&result)

	return result, nil
}
