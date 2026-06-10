package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous names of 20 location areas in the Pokemon world",
			callback:    commandMapb,
		},
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	return nil
}

func commandMap() error {
	locationAreas := getLocationArea(userInfo.MapPage)
	userInfo.MapPage++
	for _, r := range locationAreas.Results {
		fmt.Println(r.Name)
	}
	return nil
}

func commandMapb() error {
	if userInfo.MapPage == 1 {
		fmt.Println("you're on the first page")
		return nil
	}
	locationAreas := getLocationArea(userInfo.MapPage - 2)
	userInfo.MapPage--
	for _, r := range locationAreas.Results {
		fmt.Println(r.Name)
	}
	return nil
}
