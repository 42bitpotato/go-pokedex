package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/42bitpotato/go-pokedex/internal/pokeapi"
)

func commandCatch(cfg *config, arg string) error {
	url := fmt.Sprintf("%spokemon/%s/", cfg.baseUrl, arg)
	pokemon, err := pokeapi.PokemonInformation(url, cfg.cache)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	time.Sleep(2 * time.Second)
	successCatch := catchChance(pokemon.BaseExp)
	if !successCatch {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}
	fmt.Printf("%s was caught!\n", pokemon.Name)
	cfg.caughtPokemon[pokemon.Name] = pokemon
	return nil
}

func catchChance(baseExp int) bool {
	chance := (baseExp / 100) + 1
	randNr := rand.Intn(chance * 2)
	return randNr < 2
}
