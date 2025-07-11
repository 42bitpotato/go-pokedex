package main

import (
	"time"

	"github.com/42bitpotato/go-pokedex/internal/pokeapi"
	"github.com/42bitpotato/go-pokedex/internal/pokecache"
)

func main() {
	cfg := &config{
		baseUrl:       "https://pokeapi.co/api/v2/",
		mapNext:       "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
		mapPrevious:   "",
		caughtPokemon: make(map[string]pokeapi.PokemonInfo),
		cacheInterval: 5 * time.Second,
	}

	cfg.cache = pokecache.NewCache(cfg.cacheInterval)
	startRepl(cfg)
}
