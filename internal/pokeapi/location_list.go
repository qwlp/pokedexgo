package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/qwlp/pokedexgo/internal/pokecache"
)

func (c *Client) ListLocations(pageURL *string, cache *pokecache.Cache) (RespShallowLocation, error) {
	result := RespShallowLocation{}
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	item, exists := cache.Get(url)
	if exists {
		if err := json.Unmarshal(item, &result); err != nil {
			return result, err
		}
		return result, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return result, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return result, err
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}
	cache.Add(url, dat)
	if err := json.Unmarshal(dat, &result); err != nil {
		return result, err
	}

	return result, nil

}
