package pokeclient

import (
	"encoding/json"
	"fmt"
	"io"
)

type Location struct {
	Url string `json:"url"`
	Name string `json:"name"`
}

type Locations struct {
	Results []Location `json:"results"`
	Next string `json:"next"`
	Prev string `json:"previous"`
}

func (c *Client) GetLocationArea(url string) (Locations, error) {
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area/"
	}
	res, err := c.httpClient.Get(url)
	if err != nil {
		return Locations{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Locations{}, err
	}
        locations  := Locations{}
	err = json.Unmarshal(data, &locations)
	if err != nil {
		fmt.Println("Shajt: ", err)
	}

	return locations, nil
}
