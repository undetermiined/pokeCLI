package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config) error {
	locationAreas, err := cfg.pokeapiClient.GetLocationAreas(cfg.next)
	if err != nil {
		return err
	}
	cfg.page++

	fmt.Printf("----\nlocation-area: Page %d\n----\n", cfg.page)
	for _, area := range locationAreas.Results {
		fmt.Println(area.Name)
	}

	cfg.next = locationAreas.Next
	cfg.prev = locationAreas.Previous

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prev == nil {
		return errors.New("please select a page first")
	}

	locationAreas, err := cfg.pokeapiClient.GetLocationAreas(cfg.prev)
	if err != nil {
		return err
	}
	cfg.page--

	fmt.Printf("----\nlocation-area: Page %d\n----\n", cfg.page)
	for _, area := range locationAreas.Results {
		fmt.Println(area.Name)
	}

	cfg.next = locationAreas.Next
	cfg.prev = locationAreas.Previous

	return nil
}
