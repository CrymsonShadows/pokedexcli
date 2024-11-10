package main

import (
	"fmt"
	"os"

	pokeapi "github.com/CrymsonShadows/pokedexcli/internal/pokeAPI"
)

type config struct {
	pokeapiCLient   pokeapi.Client
	nextLocationURL string
	prevLocationURL string
}

type cliCommand struct {
	name        string
	description string
	callback    func(c *config, parameters ...string) error
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
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the 20 previous location areas in the Pokemon world. Each subsequent call to mapb should display the previous 20 locations.",
			callback:    commandMapPrevious,
		},
		"explore": {
			name:        "explore",
			description: "See the list of Pokemon you can encounter in the given area",
			callback:    commandExplore,
		},
	}
}

func commandHelp(c *config, parameters ...string) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n")
	fmt.Println()
	cliCommands := getCommands()
	for _, command := range cliCommands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}

func commandExit(c *config, parameters ...string) error {
	os.Exit(0)
	return nil
}

func commandMapNext(c *config, parameters ...string) error {
	if len(c.nextLocationURL) == 0 {
		fmt.Println("There are no further locations.")
		return nil
	}

	locations, err := c.pokeapiCLient.ListLocations(c.nextLocationURL)
	if err != nil {
		return fmt.Errorf("error getting location data from pokeapi: %w", err)
	}

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	c.nextLocationURL = locations.Next
	c.prevLocationURL = locations.Previous
	return nil
}

func commandMapPrevious(c *config, parameters ...string) error {
	if len(c.prevLocationURL) == 0 {
		fmt.Println("There are no locations behind.")
		return nil
	}

	locations, err := c.pokeapiCLient.ListLocations(c.prevLocationURL)
	if err != nil {
		return fmt.Errorf("error getting location data from pokeapi: %w", err)
	}

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	c.nextLocationURL = locations.Next
	c.prevLocationURL = locations.Previous
	return nil
}

func commandExplore(c *config, parameters ...string) error {
	return nil
}
