package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"

	"github.com/petomackay/pokedexcli/internal/pokeclient"
)



var mapConfig config

func getCommands() map[string]cliCommand {
    return map[string]cliCommand {
        "exit": {
            name:        "exit",
            description: "Exit the Pokedex",
            callback:    commandExit,
        },
        "help": {
            name:        "help",
            description: "Displays a help message",
            callback:    commandHelp,
        },
        "map": {
            name:        "map",
            description: "Displays the names of 20 locations in the Pokemon world.",
            callback:    commandMap,
        },
        "mapb": {
            name:        "mapb",
            description: "Displays the names of the 20 previous locations in the Pokemon world.",
            callback:    commandMapB,
        },
        "explore": {
            name:        "explore",
            description: "Lists all the pokemon located in the given area.",
            callback:    commandExplore,
        },
        "catch": {
            name:        "catch",
            description: "Tries to catch the given pokemon.",
            callback:    commandCatch,
        },
        "inspect": {
            name:        "inspect",
            description: "Show the stats for the pokemon.",
            callback:    commandInspect,
        },
        "pokedex": {
            name:        "pokedex",
            description: "List all the caught pokemon.",
            callback:    commandPokedex,
        },
    }
}

type cliCommand struct {
    name string
    description string
    callback func(*config, ...string) error
}

func commandMap(conf *config, args ...string) error {
    if conf == nil {
        return fmt.Errorf("Conf was nil")
    }
    locations, err := conf.pokeclient.GetLocationArea(conf.next)
    if err != nil {
        return err
    }

    conf.prev = locations.Prev
    conf.next = locations.Next

    for _, loc := range locations.Results {
        fmt.Printf("%s\n", loc.Name)
    }
    return nil
}

func commandMapB(conf *config, args ...string) error {
    if conf == nil {
        return fmt.Errorf("Conf was null")
    }

    if conf.prev == "" {
        fmt.Println("you're on the first page")
        return nil
    }

    locations, err := conf.pokeclient.GetLocationArea(conf.prev)
    if err != nil {
        return err
    }

    conf.next = locations.Next
    conf.prev = locations.Prev

    for _, loc := range locations.Results {
        fmt.Println(loc.Name)
    }

    return nil
}

func commandExplore(conf *config, args ...string) error {
    if conf == nil {
        return fmt.Errorf("Conf was null")
    }

    if len(args) != 1 {
        return errors.New("you must provide a location name")
    }

    areaName := args[0]

    pokemans, err := conf.pokeclient.GetLocationPokemon(areaName)
    if err != nil {
        return err
    }
    fmt.Println("Exploring " + areaName)

    fmt.Println("Found Pokemon:")
    for _, pokeman := range pokemans {
        fmt.Println(" - " + pokeman.Name)
    } 

    return nil
}

func commandCatch(conf *config, args ...string) error {
    if conf == nil {
        return errors.New("conf was null")
    }

    if len(args) != 1 {
        return errors.New("you must provide a pokemon name")
    }

    pokemonName := args[0]

    //TODO: call the API for pokemon info
    pokemon, err := conf.pokeclient.GetPokemon(pokemonName)
    if err != nil {
        return err
    }

    fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
    if tryToCatch(pokemon.Base_XP) {
        fmt.Println(pokemonName + " was caught!")
        fmt.Println("You may now inspect it with the inspect command.")
        conf.pokedex[pokemonName] = pokemon
    } else {
        fmt.Println(pokemonName + " escaped!")
    }
    
    return nil
}

func commandInspect(conf *config, args ...string) error {
    if len(args) != 1 {
        return errors.New("you must provide a pokemon name")
    }
    name := args[0]

    pokeman, found := conf.pokedex[name]
    if !found {
        fmt.Printf("You haven't caught %s yet.\n", name)
        return nil
    }
    printPokemonDetails(pokeman)
    return nil
}

func commandPokedex(conf *config, args ...string) error {
    if len(conf.pokedex) > 0 {
        fmt.Println("Your Pokedex:")
    } else {
        fmt.Println("Your Pokedex is empty. Go catch some pokemon!")
    }
    for pokemon_name, _ := range conf.pokedex {
        fmt.Println("  -", pokemon_name)
    }
    return nil
}

func commandExit(conf *config, args ...string) error {
    fmt.Println("Closing the Pokedex... Goodbye!")
    os.Exit(0)
    return nil
}

func commandHelp(conf *config, args ...string) error {
    fmt.Println("Welcome to the Pokedex!")
    fmt.Println("Usage:")
    fmt.Println()
    for _, cmd := range getCommands() {
        fmt.Printf("%s: %s\n", cmd.name, cmd.description)
    } 
    fmt.Println()
    return nil
}

func tryToCatch(base_xp int) bool {
    return rand.Intn(base_xp * 2) < 90
}

func printPokemonDetails(pokemon pokeclient.Pokemon) {
    fmt.Printf("Name: %s\nHeight: %d\nWeight:%d\n", pokemon.Name, 0, 0)
    fmt.Println("Abilities:")
    for _, a := range pokemon.Abilities {
        fmt.Printf("  - %s\n", a.Ability.Name)
    }
    fmt.Println("Moves:")
    for _, m:= range pokemon.Moves {
        fmt.Printf("  - %s\n", m.Move.Name)
    }
    fmt.Println("Stats:")
    for _, s := range pokemon.Stats {
        fmt.Printf("  - %s: %d\n", s.Stat.Name, s.Value)
    }
    fmt.Println("Types:")
    for _, t := range pokemon.Types {
        fmt.Printf("  - %s\n", t.Type.Name)
    }
}
