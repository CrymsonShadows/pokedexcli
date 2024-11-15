package main

import (
	"fmt"
	"math/rand"
	"os"

	pokeapi "github.com/CrymsonShadows/pokedexcli/internal/pokeAPI"
)

type config struct {
	pokeapiCLient   pokeapi.Client
	pokedex         map[string]pokeapi.Pokemon
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
		"catch": {
			name:        "catch",
			description: "Try to catch the given Pokemon and add them to your Pokedex",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "See the details of the given Pokemon if you have caught them",
			callback:    commandInspect,
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
	if len(parameters) != 2 {
		fmt.Println("explore command needs 1 location area to be provided")
		return nil
	}
	locationArea := parameters[1]
	locationAreaDetails, err := c.pokeapiCLient.GetLocationAreaDetails(locationArea)
	if err != nil {
		return fmt.Errorf("error getting location area details from pokeapi: %w", err)
	}

	pokeEncounters := locationAreaDetails.PokemonEncounters

	fmt.Printf("Exploring %s...\n", locationArea)
	if len(pokeEncounters) == 0 {
		fmt.Println("No Pokemon found")
	} else {
		fmt.Println("Found Pokemon:")
		for _, encounter := range pokeEncounters {
			fmt.Printf(" - %s\n", encounter.Pokemon.Name)
		}
	}
	return nil
}

func commandCatch(c *config, parameters ...string) error {
	if len(parameters) != 2 {
		fmt.Println("catch command needs 1 pokemon to be provided")
		return nil
	}
	pokemonName := parameters[1]
	pokemonDetails, err := c.pokeapiCLient.GetPokemonDetails(pokemonName)
	if err != nil {
		return fmt.Errorf("error getting pokemon details from pokeapi: %w", err)
	}

	chance := pokemonDetails.BaseExperience / 20
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonDetails.Name)
	if randInt := rand.Intn(chance); randInt == 0 {
		c.pokedex[pokemonDetails.Name] = pokemonDetails
		fmt.Printf("%s was caught!\n", pokemonDetails.Name)
	} else {
		fmt.Printf("%s escaped!\n", pokemonDetails.Name)
	}
	return nil
}

func commandInspect(c *config, parameters ...string) error {
	if len(parameters) != 2 {
		fmt.Println("inspect command needs 1 pokemon to be provided")
		return nil
	}
	pokemonName := parameters[1]
	pokemon, ok := c.pokedex[pokemonName]
	if !ok {
		fmt.Printf("You have not caught a %s yet", pokemonName)
		return nil
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("   -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, types := range pokemon.Types {
		fmt.Printf("   - %s\n", types.Type.Name)
	}
	return nil
}
