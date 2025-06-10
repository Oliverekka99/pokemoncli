package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println(">")

		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}
		commandName := cleaned[0]

		availableCommands := getCommands()
		command, exists := availableCommands[commandName]
		if !exists {
			fmt.Println("Unknown command:", commandName)
			continue
		}
		command.callback()
	}
}

type cliCommand struct {
	name				  string
	description		string
	callback			func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: "help",
			description: "Prints the help message",
			callback: callbackHelp,
		},
		"exit": {
			name: "exit",
			description: "Turn off the Pokemon REPL",
			callback: callbackExit,
		},
	}
}

func cleanInput(input string) []string {
	words := strings.Fields(strings.ToLower(input))
	if len(words) == 0 {
		return nil
	}
	return words
}