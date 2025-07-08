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

type PokemonInfo struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	BaseExp string `json:"base_experience"`
	Height  string `json:"height"`
	Weight  string `json:"weight"`
	Forms   []struct {
		Name string `json:"name"`
	} `json:"forms"`
	Species struct {
		Name string `json:"name"`
	} `json:"species"`
}
