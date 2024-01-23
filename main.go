package main

import "github.com/luluagasantiago/pokedexcli/internal/pokeapi"

// build a REPL that reads a pokemon name and prints its stats
type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
}

func main() {

	cfg := config{
		pokeapiClient:       pokeapi.NewClient(),
		nextLocationAreaURL: nil,
	}

	startRepl(&cfg)
}
