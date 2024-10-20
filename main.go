package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var cliCommands map[string]cliCommand = map[string]cliCommand{
	"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    nil,
	},
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
}

func commandHelp() error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, command := range cliCommands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandExit() error {
	return nil
}

func main() {
	if val, ok := cliCommands["help"]; ok {
		val.callback = commandHelp
		cliCommands["help"] = val
	}
	userInput := ""
	run := true
	for run {
		fmt.Print("Pokedex > ")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			userInput = scanner.Text()
			if strings.Contains(userInput, "exit") {
				run = false
				break
			}
			// output, ok := cliCommands[userInput]
			// if !ok
		}
	}
}
