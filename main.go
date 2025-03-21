package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    for true {
        fmt.Print("Pokedex > ")
        scanner.Scan()
        input := scanner.Text()
        cleanInput := cleanInput(input)
        if len(cleanInput) > 0 {
            fmt.Printf("Your command was: %s\n", cleanInput[0])
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
