package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationAreas struct {
	Count    int                   `json:"count"`
	Next     string                `json:"next"`
	Previous string                `json:"previous"`
	Results  []LocationAreasResult `json:"results"`
}

type LocationAreasResult struct {
	Name string `json:"name"`
}

func commandMap(cfg *config) error {
	limit := cfg.mapLimit
	offset := limit * cfg.mapPage

	baseUrl := "https://pokeapi.co/api/v2/location-area/"
	parameters := fmt.Sprintf("?limit=%d&offset=%d", limit, offset)

	url := fmt.Sprintf("%s%s", baseUrl, parameters)

	// Get url response
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch location areas: %w", err)
	}

	//Save response body to variable
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("failed to read location response: %w", err)
	}

	// Unmarshall json to struct
	var locationAreas []LocationAreas
	if err := json.Unmarshal(data, &locationAreas); err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	for _, area := range locationAreas[0].Results {
		fmt.Println(area.Name)
	}
	return nil
}
