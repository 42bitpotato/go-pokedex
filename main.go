package main

import (
	"github.com/42bitpotato/go-pokedex/internal/pokecache"
)

type config struct {
	mapNext       string
	mapPrevious   string
	cacheInterval int
}

func main() {
	cfg := &config{
		mapNext:       "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
		mapPrevious:   "",
		cacheInterval: 5,
	}
	pokecache.NewCache(cfg.cacheInterval)
	startRepl(cfg)
}
