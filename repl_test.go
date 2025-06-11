package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		expected := c.expected
		if len(actual) != len(expected) {
			t.Errorf("Actual list:\n%v\n\nExpected length:\n\n%v\n----\n", actual, expected)
		}
		for word := range actual {
			if actual[word] != expected[word] {
				t.Errorf("Actual word: %s\nExpected word: %s\n----\n", actual, expected)
			}
		}
	}
}
