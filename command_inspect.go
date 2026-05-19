package main

import (
	"encoding/json"
	"fmt"
)

func commandInspect(cfg *config, args []string) error {
	pokemon, exists := cfg.caughtPokemon[args[0]]
	if !exists {
		return fmt.Errorf("you have not caught that pokemon")
	}

	data, _ := json.MarshalIndent(pokemon, "", "  ")
	fmt.Println("----")
	fmt.Println(string(data))
	fmt.Println("----")

	return nil
}
