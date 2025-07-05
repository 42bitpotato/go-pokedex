package main

import "github.com/42bitpotato/go-pokedex/internal/pokecache"

func main() {
	cfg := &config{
		mapNext:       "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
		mapPrevious:   "",
		cacheInterval: 5,
	}

	cfg.cache = pokecache.NewCache(cfg.cacheInterval)
	startRepl(cfg)
}
