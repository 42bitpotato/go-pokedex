package pokeapi

type mapRespons struct {
	Count    int                   `json:"count"`
	Next     string                `json:"next"`
	Previous string                `json:"previous"`
	Results  []locationAreasResult `json:"results"`
}

type locationAreasResult struct {
	Name string `json:"name"`
}

type LocationEncounters struct {
	Name       string `json:"name"`
	Encounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}
