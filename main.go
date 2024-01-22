package main

import (
	"fmt"
	"log"

	"github.com/luluagasantiago/pokedexcli/internal/pokeapi"
)

//build a REPL that reads a pokemon name and prints its stats

func main() {
	pokeapiClient := pokeapi.NewClient()

	resp, err := pokeapiClient.ListLocationAreas()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)

	//startRepl()

}
