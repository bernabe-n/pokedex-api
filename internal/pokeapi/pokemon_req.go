package pokeapi

import (
	"net/http";
	"fmt";
	"io";
	"encoding/json"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	endpoint := "/pokemon/" + pokemonName
	fullURL := baseURL + endpoint

	data, ok :=c.cache.Get(fullURL)
	if ok {
		fmt.Printf("cache hit for %s\n", fullURL)
		pokemon := Pokemon{}
		err := json.Unmarshal(data, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}

		return pokemon, nil
	}
	fmt.Printf("cache miss for %s\n", fullURL)

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)//executes the request
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(fullURL, data)

	return pokemon, nil
}