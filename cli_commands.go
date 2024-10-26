package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	pokeapi "github.com/CrymsonShadows/pokedexcli/internal/pokeAPI"
)

type config struct {
	Next     string
	Previous string
}

type cliCommand struct {
	name        string
	description string
	callback    func(c *config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world. Each subsequent call to map should display the next 20 locations.",
			callback:    commandMapNext,
		},
	}
}

func commandHelp(c *config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n")
	fmt.Println()
	cliCommands := getCommands()
	for _, command := range cliCommands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}

func commandExit(c *config) error {
	os.Exit(0)
	return nil
}

func commandMapNext(c *config) error {
	if len(c.Next) == 0 {
		return errors.New("there are no further locations")
	}
	locData, err := pokeapi.GetPokeData(c.Next)
	if err != nil {
		return fmt.Errorf("error getting location data from pokeapi: %w", err)
	}
	locations := pokeapi.Location{}
	err = json.Unmarshal(locData, &locations)
	if err != nil {
		return fmt.Errorf("error unmarshalling pokeapi location json: %w", err)
	}
	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	c.Next = locations.Next
	c.Previous = locations.Previous
	return nil
}
