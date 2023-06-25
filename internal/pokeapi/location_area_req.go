package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	// if url passed then use that url else use base url
	if pageURL != nil {
		fullURL = *pageURL
	}

	// check the cache
	data, ok := c.cache.Get(fullURL)
	if ok {
		// cache hit
		fmt.Println("cache hit!")
		// unmarshall the json response into a struct
		locationAreasResp := LocationAreasResp{}
		err := json.Unmarshal(data, &locationAreasResp)
		if err != nil {
			return LocationAreasResp{}, err
		}
		return locationAreasResp, nil
	}
	// cache miss
	fmt.Println("cache miss")
	// create a new http request
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResp{}, nil
	}
	// execute the request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, nil
	}
	// close the response just before returning from the function
	defer resp.Body.Close()
	// check the status code
	if resp.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}
	// read all data
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}
	// unmarshall the json response into a struct
	locationAreasResp := LocationAreasResp{}
	err = json.Unmarshal(data, &locationAreasResp)
	if err != nil {
		return LocationAreasResp{}, err
	}

	// add data to cache
	c.cache.Add(fullURL, data)

	return locationAreasResp, nil
}
