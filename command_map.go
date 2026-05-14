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
	cfg.page++

	fmt.Printf("----\nlocation-area: Page %d\n----\n", cfg.page)
	for _, area := range locationAreas.Results {
		fmt.Println(area.Name)
	}

	if locationAreas.Next != nil {
		cfg.next = *locationAreas.Next
	} else {
		cfg.next = ""
	}
	if locationAreas.Previous != nil {
		cfg.prev = *locationAreas.Previous
	} else {
		cfg.prev = ""
	}

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prev == "" {
		return errors.New("please select a page first")
	}

	pageURL := cfg.prev

	locationAreas, err := pokeapi.GetLocationAreas(pageURL)
	if err != nil {
		return err
	}
	cfg.page--

	fmt.Printf("----\nlocation-area: Page %d\n----\n", cfg.page)
	for _, area := range locationAreas.Results {
		fmt.Println(area.Name)
	}

	if locationAreas.Next != nil {
		cfg.next = *locationAreas.Next
	} else {
		cfg.next = ""
	}
	if locationAreas.Previous != nil {
		cfg.prev = *locationAreas.Previous
	} else {
		cfg.prev = ""
	}

	return nil
}
