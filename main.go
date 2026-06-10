package main

import (
	"time"

	"github.com/qwlp/pokedexgo/internal/pokeapi"
	"github.com/qwlp/pokedexgo/internal/pokecache"
)

func main() {
	const baseTime = 5 * time.Minute
	const waitTime = baseTime + 5*time.Millisecond

	pokeClient := pokeapi.NewClient(5 * time.Second)
	cache := pokecache.NewCache(waitTime)
	cfg := &Config{
		pokeapiClient: pokeClient,
		Cache:         cache,
	}

	startREPL(cfg)
}
