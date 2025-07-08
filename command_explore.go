package main

import (
	"fmt"

	"github.com/42bitpotato/go-pokedex/internal/pokeapi"
)

func commandExplore(cfg *config, arg string) error {
	url := fmt.Sprintf("%slocation-area/%s/", cfg.baseUrl, arg)
	encountersJson, err := pokeapi.ListEncounters(url, cfg.cache)
	if err != nil {
		return err
	}
	encounters := encountersJson.Encounters
	for _, val := range encounters {
		name := val.Pokemon.Name
		fmt.Println(name)
	}
	return nil
}
