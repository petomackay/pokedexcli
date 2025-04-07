package main

import (
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
    }
}

type cliCommand struct {
    name string
    description string
    callback func(*config) error
}

func commandMap(conf *config) error {
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

func commandMapB(conf *config) error {
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

func commandExit(conf *config) error {
    fmt.Println("Closing the Pokedex... Goodbye!")
    os.Exit(0)
    return nil
}

func commandHelp(conf *config) error {
    fmt.Println("Welcome to the Pokedex!")
    fmt.Println("Usage:")
    fmt.Println()
    for _, cmd := range getCommands() {
        fmt.Printf("%s: %s\n", cmd.name, cmd.description)
    } 
    fmt.Println()
    return nil
}
