package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)

	re, err := regexp.Compile(`\s+`)
	if err != nil {
		fmt.Errorf("Invalid regex: %v", err)
	}
	words := re.Split(text, -1)

	return words
}
