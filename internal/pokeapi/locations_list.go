package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ListLocations(url string) (mapRespons, error) {
	// Get url response
	res, err := http.Get(url)
	if err != nil {
		return mapRespons{}, fmt.Errorf("failed to fetch location areas: %w", err)
	}

	//Save response body to variable
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return mapRespons{}, fmt.Errorf("failed to read location response: %w", err)
	}

	// Unmarshall json to struct
	var responsJson mapRespons
	if err := json.Unmarshal(data, &responsJson); err != nil {
		return mapRespons{}, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return responsJson, nil
}
