package main

import (
	"fmt"
	"math/rand"

	"github.com/42bitpotato/go-pokedex/internal/pokeapi"
)

func commandCatch(cfg *config, arg string) error {
	url := fmt.Sprintf("%spokemon/%s/", cfg.baseUrl, arg)
	pokemon, err := pokeapi.PokemonInformation(url, cfg.cache)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...", pokemon.Name)
	successCatch := catchChance(pokemon.BaseExp)
	if !successCatch {
		fmt.Printf("%s escaped!", pokemon.Name)
		return nil
	}

	return nil
}

func catchChance(baseExp int) bool {
	randNr := rand.Intn(baseExp)
	successLimit := (baseExp * 75) / 100
	if randNr > successLimit {
		return false
	}
	return true
}
