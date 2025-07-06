package pokeapi

import (
	"testing"
	"time"

	"github.com/42bitpotato/go-pokedex/internal/pokecache"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected [2]string
	}{
		{
			input:    "https://pokeapi.co/api/v2/location",
			expected: [2]string{"canalave-city", "https://pokeapi.co/api/v2/location?offset=20&limit=20"},
		},
		{
			input:    "https://pokeapi.co/api/v2/location?offset=20&limit=20",
			expected: [2]string{"ruin-maniac-cave", "https://pokeapi.co/api/v2/location?offset=40&limit=20"},
		},
	}
	cache := pokecache.NewCache(10 * time.Second)
	for _, c := range cases {
		result, err := ListLocations(c.input, cache)
		if err != nil {
			t.Errorf("Error running case:\n%v\n\nError: %v", c, err)
		}
		expected := c.expected
		if result.Next != expected[1] {
			t.Errorf("Actual next-url:\n%v\n\nExpected url:\n\n%v\n----\n", result.Next, expected[1])
		}
		locName := result.Results[0].Name
		if locName != expected[0] {
			t.Errorf("Actual result name:\n%v\n\nExpected name:\n\n%v\n----\n", locName, expected[0])
		}
		resultLength := len(result.Results)
		if resultLength != 20 {
			t.Errorf("Actual length: %v\nExpected name: 20\n----\n", resultLength)
		}
	}
}
