package main

import (
	"time"

	"github.com/undetermiined/pokeCLI/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 2*time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
		caughtPokemon: map[string]pokeapi.PokemonResponse{},
	}

	startRepl(cfg)
}
