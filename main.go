package main

import (
    "fmt"
    "strings"
)

func main() {
    fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
    if text == "" {
        return []string{}
    }
    trimmed := strings.TrimSpace(text) 
    return strings.Fields(strings.ToLower(trimmed))
}
