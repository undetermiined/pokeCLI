package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args []string) error {
	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("Your Pokedex is empty!")
		return nil
	}

	fmt.Println("----")
	fmt.Print("Your Pokedex:")
	for name, _ := range cfg.caughtPokemon {
		fmt.Printf("\n - %s", name)
	}
	fmt.Println("\n----")
	return nil
}
