package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	endpoint := "/location-area/"
	fullURL := baseURL + endpoint + pokemonName

	//check the cache
	dat, ok := c.cache.Get(fullURL)
	if ok {
		//cache hit
		fmt.Println("cache hit!")
		pokemonResp := Pokemon{}

		err := json.Unmarshal(dat, &pokemonResp)

		if err != nil {
			return Pokemon{}, err
		}
		return pokemonResp, nil
	}
	fmt.Println("cache miss!")
	req, err := http.NewRequest("GET", fullURL, nil)

	if err != nil {
		return Pokemon{}, err
	}
	resp, err := c.httpClient.Do(req)

	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("bad status code: %v", resp.StatusCode)

	}

	dat, err = io.ReadAll(resp.Body)

	if err != nil {
		return Pokemon{}, err
	}

	pokemonResp := Pokemon{}

	err = json.Unmarshal(dat, &pokemonResp)

	if err != nil {
		return Pokemon{}, err
	}
	c.cache.Add(fullURL, dat)

	return pokemonResp, nil

}
