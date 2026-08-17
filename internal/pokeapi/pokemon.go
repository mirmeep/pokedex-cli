package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(url string) (Pokemon, error) {
	cacheData, exists := c.cache.Get(url)
	var response Pokemon
	if exists {
		if err := json.Unmarshal(cacheData, &response); err != nil {
			return Pokemon{}, err
		}
	} else {
		res, err := http.Get(url)
		if err != nil {
			return Pokemon{}, err
		}
		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
		if res.StatusCode > 299 {
			return Pokemon{}, fmt.Errorf("Pokemon not found")
		}
		if err != nil {
			return Pokemon{}, err
		}

		c.cache.Add(url, data)

		if err := json.Unmarshal(data, &response); err != nil {
			return Pokemon{}, err
		}
	}

	return response, nil
}
