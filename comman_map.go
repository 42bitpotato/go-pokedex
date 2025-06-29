package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type mapRespons struct {
	Count    int                   `json:"count"`
	Next     string                `json:"next"`
	Previous string                `json:"previous"`
	Results  []locationAreasResult `json:"results"`
}

type locationAreasResult struct {
	Name string `json:"name"`
}

func commandMap(cfg *config) error {
	url := cfg.mapNext

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
	var mapRespons []mapRespons
	if err := json.Unmarshal(data, &mapRespons); err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}
	respons := mapRespons[0]
	for _, area := range respons.Results {
		fmt.Println(area.Name)
	}
	cfg.mapNext = respons.Next
	cfg.mapPrevious = respons.Previous

	return nil
}
