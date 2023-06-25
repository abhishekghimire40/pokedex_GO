package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	endpoint := "/pokemon/" + pokemonName
	fullURL := baseURL + endpoint

	// check the cache
	data, ok := c.cache.Get(fullURL)
	if ok {
		// cache hit
		fmt.Println("cache hit!")
		// unmarshall the json response into a struct
		pokemon := Pokemon{}
		err := json.Unmarshal(data, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}
	// cache miss
	fmt.Println("cache miss")
	// create a new http request
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Pokemon{}, nil
	}
	// execute the request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, nil
	}
	// close the response just before returning from the function
	defer resp.Body.Close()
	// check the status code
	if resp.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}
	// read all data
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}
	// unmarshall the json response into a struct
	pokemon := Pokemon{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	// add data to cache
	c.cache.Add(fullURL, data)

	return pokemon, nil
}
