package pokeapi

import (
	"fmt"
	"net/http"
	"io"
	"encoding/json"
)

func (c *Client) GetEncounters(url string) (RespEncounters, error) {
	cacheData, exists := c.cache.Get(url)
	var response RespEncounters
	if exists {
		if err := json.Unmarshal(cacheData, &response); err != nil {
			return RespEncounters{}, err
		}
	} else {
		res, err := http.Get(url)
		if err != nil {
			return RespEncounters{}, err
		}
		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
		if res.StatusCode > 299 {
			return RespEncounters{}, fmt.Errorf("Location area not found")
		}
		if err != nil {
			return RespEncounters{}, err
		}

		c.cache.Add(url, data)

		if err := json.Unmarshal(data, &response); err != nil {
			return RespEncounters{}, err
		}
	}
	return response, nil
}