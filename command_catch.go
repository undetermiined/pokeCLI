package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func commandCatch(cfg *config, args []string) error {
	if len(args) != 1 {
		return errors.New("please input single Pokemon name")
	}

	_, exists := cfg.caughtPokemon[args[0]]
	if exists {
		return fmt.Errorf("you have already caught %s", args[0])
	}

	pokeResp, err := cfg.pokeapiClient.GetPokemon(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("----\nThrowing a Pokeball at wild %s", args[0])
	for range 3 {
		time.Sleep(300 * time.Millisecond)
		fmt.Print(".")
		os.Stdout.Sync() // force flush so the dot shows immediately
	}

	chance := 100 - pokeResp.BaseExperience/2
	roll := rand.Intn(75)
	if chance < 15 {
		chance = 15
	}

	if roll <= chance {
		fmt.Printf("\nYou caught a wild %s!\nIt may now be inspected with the inspect command.\n", args[0])
		cfg.caughtPokemon[args[0]] = pokeResp
	} else {
		fmt.Printf("\nThe wild %s fled!\n", args[0])
	}
	fmt.Println("----")

	return nil
}
