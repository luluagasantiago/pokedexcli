package main

import (
	"time"

	"github.com/luluagasantiago/pokedexcli/internal/pokeapi"
)

// build a REPL that reads a pokemon name and prints its stats
type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
	caughtPokemon       map[string]pokeapi.Pokemon
}

func main() {

	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}

	startRepl(&cfg)
}
