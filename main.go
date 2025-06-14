package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
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
