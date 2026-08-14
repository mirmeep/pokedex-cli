package pokeapi

import (
	"fmt"
	"net/http"
	"io"
	"encoding/json"
)

func (c *Client) GetMapLocationAreas(url string) (RespLocations, error) {
	cacheData, exists := c.cache.Get(url)
	var response RespLocations
	if exists {
		if err := json.Unmarshal(cacheData, &response); err != nil {
			return RespLocations{}, err
		}
	} else {
		res, err := http.Get(url)
		if err != nil {
			return RespLocations{}, err
		}
		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
		if res.StatusCode > 299 {
			return RespLocations{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, res.Body)
		}
		if err != nil {
			return RespLocations{}, err
		}

		c.cache.Add(url, data)

		if err := json.Unmarshal(data, &response); err != nil {
			return RespLocations{}, err
		}
	}
	return response, nil
}