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
	resp, err := http.Get(pageUrl)
	if err != nil {
		return LocationAreasResponse{}, fmt.Errorf("location-area request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return LocationAreasResponse{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return LocationAreasResponse{}, fmt.Errorf("decoding location-area response failed: %w", err)
	}

	return result, nil
}
