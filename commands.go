package main

import (
	"errors"
	"fmt"
	"os"
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
            name:        "map",
            description: "Displays the names of 20 locations in the Pokemon world.",
            callback:    commandMapB,
        },
        "explore": {
            name:       "explore",
            description: "Lists all the pokemon located in the given area.",
            callback:   commandExplore,
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
