package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/petomackay/pokedexcli/internal/pokeclient"
)

type config struct {
    prev string
    next string
    pokeclient pokeclient.Client
    pokedex map[string]pokeclient.Pokemon
}

func main() {
    conf := config{
        pokeclient: pokeclient.NewClient(time.Second * 5, time.Minute * 5),
        pokedex:    make(map[string]pokeclient.Pokemon),
    }
    
    scanner := bufio.NewScanner(os.Stdin)
    for {
        fmt.Print("Pokedex > ")
        scanner.Scan()
        input := scanner.Text()
        cleanInput := cleanInput(input)
        if len(cleanInput) == 0 {
            continue
        }

        cmd := cleanInput[0]
        args := []string{}
        if len(cleanInput) > 1 {
            args = cleanInput[1:]
        }
        if cmd, exists := getCommands()[cmd]; exists {
            err := cmd.callback(&conf, args...)
            if err != nil {
                fmt.Print(err)
            }
        } else {
            fmt.Println("Unknown command")
        }
    }
}

func cleanInput(text string) []string {
    if text == "" {
        return []string{}
    }
    trimmed := strings.TrimSpace(text) 
    return strings.Fields(strings.ToLower(trimmed))
}
