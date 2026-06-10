package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type LocationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// helper functions

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	words := strings.Fields(lower)

	return words
}

func httpReq[T any](method string, url string) (T, error) {
	var result T
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return result, err
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return result, err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&result); err != nil {
		return result, err
	}

	return result, nil
}

func getLocationArea(page int) LocationArea {
	offset := page * 20
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?offset=%d&limit=20", offset)
	res, err := httpReq[LocationArea]("GET", url)
	if err != nil {
		fmt.Println("Something went wrong: ", err)
	}

	return res
}
