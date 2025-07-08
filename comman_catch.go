package main

import (
	"fmt"

	"github.com/42bitpotato/go-pokedex/internal/pokeapi"
)

func commandCatch(cfg *config, arg string) error {
	url := fmt.Sprintf("%slocation-area/%s/", cfg.baseUrl, arg)
	pokemon, err := pokeapi.ListPokemon(url, cfg.cache)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...", pokemon)

	return nil
}
