package pokeapi

import (
	"net/http";
	"fmt";
	"io";
	"encoding/json"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreaResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

	if pageURL != nil {
		fullURL = *pageURL
	}

	data, ok :=c.cache.Get(fullURL)
	if ok {
		fmt.Printf("cache hit for %s\n", fullURL)
		locationAreaResp := LocationAreaResp{}
		err := json.Unmarshal(data, &locationAreaResp)
		if err != nil {
		return LocationAreaResp{}, err
		}

		return locationAreaResp, nil
	}
	fmt.Printf("cache miss for %s\n", fullURL)

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreaResp{}, err
	}

	resp, err := c.httpClient.Do(req)//executes the request
	if err != nil {
		return LocationAreaResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreaResp{}, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResp{}, err
	}

	locationAreaResp := LocationAreaResp{}
	err = json.Unmarshal(data, &locationAreaResp)
	if err != nil {
		return LocationAreaResp{}, err
	}

	c.cache.Add(fullURL, data)

	return locationAreaResp, nil
}