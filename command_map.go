package main

import (
	"errors"
	"fmt"

	"github.com/undetermiined/pokeCLI/internal/pokeapi"
)

func commandMap(cfg *config) error {
	pageURL := "https://pokeapi.co/api/v2/location-area"

	if cfg.next != "" {
		pageURL = cfg.next
	}

	locationAreas, err := pokeapi.GetLocationAreas(pageURL)
	if err != nil {
		return err
	}

	for _, area := range locationAreas.Results {
		fmt.Println(area.Name)
	}

	*cfg.next = *locationAreas.Next
	*cfg.prev = *locationAreas.Previous

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prev == nil {
		return errors.New("Please select a page first!")
	}

	locationAreas, err := pokeapi.GetLocationAreas(*cfg.prev)
	if err != nil {
		return err
	}

	for _, area := range locationAreas.Results {
		fmt.Println(area.Name)
	}

	*cfg.next = *locationAreas.Next
	*cfg.prev = *locationAreas.Previous

	return nil
}
