package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/undetermiined/pokeCLI/internal/pokeapi"
)

func startRepl(cfg *config) {
	// Begin listening
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")

		//	Break/exit at EOF or error
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "error:", err)
				os.Exit(1)
			}
			break
		}

		words := cleanInput(scanner.Text())

		// Skip at empty index
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

type config struct {
	pokeapiClient pokeapi.Client
	next          *string
	prev          *string
	page          int
}

type cliCommand struct {
	name     string
	desc     string
	callback func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:     "exit",
			desc:     "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name:     "help",
			desc:     "Displays a help message",
			callback: commandHelp,
		},
		"map": {
			name:     "map",
			desc:     "Display a list of 20 location-area pairs: Display next 20 items on recurrence",
			callback: commandMap,
		},
		"mapb": {
			name:     "mapb",
			desc:     "Display previous location-area pair list",
			callback: commandMapb,
		},
	}
}
