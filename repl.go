package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()

		wordsList, err := cleanInput(text)
		if err != nil {
			fmt.Printf("error formatting input: %v\n", err)
			continue
		}
		firstWord := wordsList[0]

		command, ok := getCommands()[firstWord]
		if !ok {
			fmt.Print("Unknown command\n")
			continue
		}
		err = command.callback(cfg)
		if err != nil {
			fmt.Printf("error while running command '%s': %v\n", firstWord, err)
		}
	}
}

func cleanInput(text string) ([]string, error) {
	text = strings.TrimSpace(text)

	// Check if only whitespaces
	if text == "" {
		return nil, fmt.Errorf("input is blank or only whitespace")
	}
	// Check if non-alphanumeric
	reAlpha := regexp.MustCompile(`^[A-Za-z\s]+$`)
	if !reAlpha.MatchString(text) {
		return nil, fmt.Errorf("input containing non-alphanumeric characters: %s", text)
	}

	text = strings.ToLower(text)

	re, err := regexp.Compile(`\s+`)
	if err != nil {
		return nil, fmt.Errorf("invalid regex: %w", err)
	}
	words := re.Split(text, -1)

	return words, nil
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
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
			description: "Display the nextr page with names of 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display previous page with names of 20 location areas in the Pokemon world",
			callback:    commandMapB,
		},
	}
}
