package main

import (
	"fmt"
	"strings"
)

func commandInspect(cfg *config, args []string) error {
	pokemon, exists := cfg.caughtPokemon[args[0]]
	if !exists {
		return fmt.Errorf("you have not caught that pokemon")
	}

	fmt.Println("----")
	fmt.Printf("Dex #: %d\n", pokemon.Order)
	fmt.Printf("Name: %s\n", strings.ToUpper(pokemon.Name[:1])+pokemon.Name[1:])
	fmt.Printf("Height: %d cm\n", pokemon.Height*10)
	fmt.Printf("Weight: %.2f lbs\n", float64(pokemon.Weight)/4.536)
	fmt.Printf("Base Exp: %d\n", pokemon.BaseExperience)
	fmt.Print("Stats:")
	for _, stats := range pokemon.Stats {
		fmt.Printf("\n -%s: %d", strings.ToUpper(stats.Stat.Name[:1])+stats.Stat.Name[1:], stats.BaseStat)
	}
	fmt.Print("\nTypes:")
	for _, typ := range pokemon.Types {
		fmt.Printf("\n - %s", strings.ToUpper(typ.Type.Name[:1])+typ.Type.Name[1:])
	}
	fmt.Print("\nForms:")
	for _, form := range pokemon.Forms {
		fmt.Printf("\n - %s", strings.ToUpper(form.Name[:1])+form.Name[1:])
	}
	fmt.Println("\n----")

	return nil
}
