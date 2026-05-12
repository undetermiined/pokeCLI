package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	//	pokeCLI starts listening
	for {
		fmt.Print("Pokedex > ")

		//	Break at EOF or error
		if !scanner.Scan() {
			break
		}

		line := scanner.Text()
		words := cleanInput(line)

		// Skip at empty index
		if len(words) == 0 {
			continue
		}
		commandName := words[0]

		command, ok := commandRegistry[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		if err := command.callback(); err != nil {
			fmt.Println(err)
		}
	}

	// Final error check
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
