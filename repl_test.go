package main

import (
	"regexp"
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
		actual, err := cleanInput(c.input)
		if err != nil {
			t.Errorf("Error running case:\n%v\n\nError: %v", c, err)
		}

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

func TestCleanInputErrors(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " ",
			expected: nil,
		},
		{
			input:    "...",
			expected: nil,
		},
		{
			input:    "pikachu, charmander, some other pokemon",
			expected: nil,
		},
		{
			input:    "pika 2 cent",
			expected: nil,
		},
	}
	for _, c := range cases {
		re := regexp.MustCompile(`^[A-Za-z]+$`)
		if re.MatchString(c.input) {
			t.Errorf("input value %s, want match for anything not A-Z a-z", c.input)
		}
		actual, err := cleanInput(c.input)
		if err == nil {
			t.Errorf("actual output:\n%v\n\nexpected error\n----\n", actual)
		}
	}
}
