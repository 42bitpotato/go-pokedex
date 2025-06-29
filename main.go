package main

type config struct {
	mapNext     string
	mapPrevious string
}

func main() {
	cfg := &config{
		mapNext:     "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
		mapPrevious: "",
	}
	startRepl(cfg)
}
