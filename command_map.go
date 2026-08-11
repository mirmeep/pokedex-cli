package main

import (
	"fmt"
	"net/http"
	"io"
	"encoding/json"
	"errors"
)

type Response struct {
	Results []Location `json:"results"`
	Next *string
	Previous *string

}

type Location struct {
	Name string `json:"name"`
}

func getMapLocationAreas(url string) (Response, error) {
	res, err := http.Get(url)
	if err != nil {
		return Response{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return Response{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, res.Body)
	}
	if err != nil {
		return Response{}, err
	}

	var response Response
	if err := json.Unmarshal(data, &response); err != nil {
		return Response{}, err
	}

	return response, nil
}

func mapResponseHandler(baseURL string, c *config) error {
	response, err := getMapLocationAreas(baseURL)
	if err != nil {
		return err
	}
	c.next = response.Next
	c.previous = response.Previous
	for _, location := range response.Results {
		fmt.Println(location.Name)
	}
	
	return nil
}

func commandMapB(c *config) error {
	baseURL := "https://pokeapi.co/api/v2/location-area/"
	if c.previous == nil {
		return errors.New("on the first page, can't go back")
	}	
	baseURL = *c.previous
	mapResponseHandler(baseURL, c)

	return nil
}

func commandMap(c *config) error {
	baseURL := "https://pokeapi.co/api/v2/location-area/"
	if c.next != nil {
		baseURL = *c.next
	}

	mapResponseHandler(baseURL, c)

	return nil
}