package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
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
        if cmd, exists := getCommands()[cmd]; exists {
            err := cmd.callback()
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
