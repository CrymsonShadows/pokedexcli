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
		command.callback()
	}
}
