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

func runRepl(c *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
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
		err := command.callback(c, cleanedInput...)
		if err != nil {
			fmt.Printf("%v", err)
			return
		}
	}
}
