package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(input string) []string {
	lowered := strings.ToLower(input)
	return strings.Fields(lowered)
}

func runRepl() {
	c := config{
		Next:     "https://pokeapi.co/api/v2/location-area",
		Previous: "",
	}
	for {
		fmt.Print("Pokedex > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		userInput := scanner.Text()
		cleanedInput := cleanInput(userInput)
		if len(cleanedInput) == 0 {
			continue
		}

		cliCommands := getCommands()
		userCommand := cleanedInput[0]
		command, ok := cliCommands[userCommand]
		if !ok {
			fmt.Println(userCommand, "is not a command")
			continue
		}
		err := command.callback(&c)
		if err != nil {
			fmt.Printf("%v", err)
			return
		}
	}
}
