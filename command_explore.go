package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func commandExplore(cfg *config, args []string) error {
	if len(args) != 1 {
		return errors.New("please input single hyphen sensitive location-area name")
	}

	pokeEncounters, err := cfg.pokeapiClient.GetAreaPokemon(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("----\nExploring %v", args[0])
	for range 3 {
		time.Sleep(100 * time.Millisecond)
		fmt.Print(".")
		os.Stdout.Sync() // force flush so the dot shows immediately
	}
	fmt.Println("\nFound Pokemon!:")
	for _, poke := range pokeEncounters.PokemonEncounters {
		name := "- " + poke.Pokemon.Name
		time.Sleep(50 * time.Millisecond)
		fmt.Println(name)
	}
	fmt.Println("---")

	return nil
}
