package main

import (
	"bufio"
	"fmt"
	"os"
)

type UserInfo struct {
	MapPage int
}

var userInfo UserInfo

func init() {
	userInfo = UserInfo{
		MapPage: 0,
	}
}

func executeCallback(input string) error {
	cmd, exists := getCommands()[input]
	if !exists {
		fmt.Println("Unknown command")
	}

	err := cmd.callback()
	if err != nil {
		return err
	}
	return nil
}

// entry point for REPL

func startREPL() {

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
		executeCallback(command)
	}
}
