package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/qwlp/pokedexgo/internal/pokeapi"
	"github.com/qwlp/pokedexgo/internal/pokecache"
)

type Config struct {
	pokeapiClient       pokeapi.Client
	nextLocationURL     *string
	previousLocationURL *string
	Cache               *pokecache.Cache
}

// entry point for REPL

func startREPL(cfg *Config) {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()

		if len(input) == 0 {
			continue
		}

		cleanedInput := cleanInput(input)

		command := cleanedInput[0]
		cmd, exists := getCommands()[command]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		err := cmd.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Provides the help menu",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous names of 20 location areas in the Pokemon world",
			callback:    commandMapb,
		},
	}
}
