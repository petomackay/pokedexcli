package pokeclient

import (
	"encoding/json"
	"io"
)

const (
    baseUrl = "https://pokeapi.co/api/v2"
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

// jq .pokemon_encounters.[].pokemon.name
type LocationAreaDetails struct {
	PokemonEncounters []struct {
		PokemonEncounter PokemonEncounter `json:"pokemon"`
	} `json:"pokemon_encounters"`
}
type PokemonEncounter struct {
	Name string `json:"name"`
}

type Pokemon struct {
	Name string `json:"name"`
	Base_XP int `json:"base_experience"`
}

func (c *Client) GetLocationArea(url string) (Locations, error) {
	if url == "" {
		url = baseUrl + "/location-area/?offset=0&limit=20"
	}

	var data []byte
        if value, cached := c.cache.Get(url); cached {
		data = value
	} else {
                res, err := c.httpClient.Get(url)
        	if err != nil {
        		return Locations{}, err
        	}
        	defer res.Body.Close()
        
        	data, err = io.ReadAll(res.Body)
        	if err != nil {
        		return Locations{}, err
        	}
		c.cache.Add(url, data)
	}

        locations  := Locations{}
	err := json.Unmarshal(data, &locations)
	if err != nil {
		return Locations{}, err
	}

	return locations, nil
}

func (c *Client) GetLocationPokemon(location_name string) ([]PokemonEncounter, error) {
	url := baseUrl + "/location-area/" + location_name
	var data []byte
	if value, cached := c.cache.Get(url); cached {
		data = value
	} else {
		res, err := c.httpClient.Get(url)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		c.cache.Add(url, data)
	}

	areaDetails := LocationAreaDetails{}
	err := json.Unmarshal(data, &areaDetails)
	if err != nil {
		return nil, err
	}
	pokemon := make([]PokemonEncounter, 0, len(areaDetails.PokemonEncounters))
	for _, encounter := range areaDetails.PokemonEncounters {
		pokemon = append(pokemon, encounter.PokemonEncounter)
	}
	return pokemon, err

}

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseUrl + "/pokemon/" + name
	var data []byte
	if value, cached := c.cache.Get(url); cached {
		data = value
	} else {
		res, err := c.httpClient.Get(url)
		if err != nil {
			return Pokemon{}, err
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return Pokemon{}, err
		}
		c.cache.Add(url, data)
	}

	pokemon := Pokemon{}
	err := json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, nil
	}
	return pokemon, nil
}
