package main

import (
	"time"

	pokeapi "github.com/CrymsonShadows/pokedexcli/internal/pokeAPI"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)

	c := &config{
		pokeapiCLient:   pokeClient,
		nextLocationURL: "https://pokeapi.co/api/v2/location-area",
		pokedex:         make(map[string]pokeapi.Pokemon),
	}

	runRepl(c)
}
