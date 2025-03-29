package main

import "fmt"
import "os"

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
    }
}

type cliCommand struct {
    name string
    description string
    callback func() error
}

func commandExit() error {
    fmt.Println("Closing the Pokedex... Goodbye!")
    os.Exit(0)
    return nil
}

func commandHelp() error {
    fmt.Println("Welcome to the Pokedex!")
    fmt.Println("Usage:\n")
    fmt.Println()
    for _, cmd := range getCommands() {
        fmt.Printf("%s: %s\n", cmd.name, cmd.description)
    } 
    fmt.Println()
    return nil
}
