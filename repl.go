package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var cliCommands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
}

func startRepl() {
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

		command, ok := cliCommands[firstWord]
		if !ok {
			fmt.Print("Unknown command\n")
			continue
		}
		err = command.callback()
		if err != nil {
			fmt.Printf("error while running command '%s': %v", firstWord, err)
		}
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
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
