package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/42bitpotato/go-pokedex/internal/pokecache"
)

func ListLocations(url string, cache *pokecache.Cache) (mapRespons, error) {

	// Check if url in cache
	data, ok := cache.Get(url)
	if !ok {
		// Get url response
		res, err := http.Get(url)
		if err != nil {
			return mapRespons{}, fmt.Errorf("failed to fetch location areas: %w", err)
		}

		//Save response body to variable
		defer res.Body.Close()
		data, err = io.ReadAll(res.Body)
		if err != nil {
			return mapRespons{}, fmt.Errorf("failed to read location response: %w", err)
		}

	}

	// Unmarshall json to struct
	var responsJson mapRespons
	if err := json.Unmarshal(data, &responsJson); err != nil {
		return mapRespons{}, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	cache.Add(url, data)
	return responsJson, nil
}
