package main

import (
	"fmt"

	"github.com/42bitpotato/go-pokedex/internal/pokeapi"
)

func commandMap(cfg *config) error {
	url := cfg.mapNext
	areas, err := pokeapi.ListLocations(url)
	if err != nil {
		return err
	}

	for _, area := range areas.Results {
		fmt.Println(area.Name)
	}

	// Update config struct
	cfg.mapNext = areas.Next
	cfg.mapPrevious = areas.Previous

	return nil
}

func commandMapB(cfg *config) error {
	if cfg.mapPrevious == "" {
		fmt.Println("You are on first page, cant go back.")
		return nil
	}
	url := cfg.mapPrevious
	areas, err := pokeapi.ListLocations(url)
	if err != nil {
		return err
	}

	for _, area := range areas.Results {
		fmt.Println(area.Name)
	}

	// Update config struct
	cfg.mapNext = areas.Next
	cfg.mapPrevious = areas.Previous

	return nil
}
