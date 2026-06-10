package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *Config) error {
	locationAreas, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationURL, cfg.Cache)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = locationAreas.Next
	cfg.previousLocationURL = locationAreas.Previous

	for _, r := range locationAreas.Results {
		fmt.Println(r.Name)
	}
	return nil
}

func commandMapb(cfg *Config) error {
	if cfg.previousLocationURL == nil {
		return errors.New("you're on the first page")
	}
	locationAreas, err := cfg.pokeapiClient.ListLocations(cfg.previousLocationURL, cfg.Cache)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = locationAreas.Next
	cfg.previousLocationURL = locationAreas.Previous

	for _, r := range locationAreas.Results {
		fmt.Println(r.Name)
	}
	return nil
}
