package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
    //var nilString string
    //var nilSlice []string
    cases := []struct {
        input string
        expected []string
    } {
        {
            input: " hello world ",
            expected: []string{"hello", "world"},
        },
        {
            input: "",
            expected: []string{},
        },
        {
            input: "  wHy  are  You    tYping Like THIS???  ",
            expected: []string{"why", "are", "you", "typing", "like", "this???"},
        },
        {
		input:    "  ",
		expected: []string{},
	},
    }


    for _,c := range cases {
        actual := cleanInput(c.input)
        if len(actual) != len(c.expected) {
            t.Errorf("Wrong length of the returned slice for input: '%s'\nExpected: %d Was: %d", c.input, len(c.expected), len(actual))
        }
        for i, token := range actual {
            if token != c.expected[i] {
                t.Errorf("Expected: %s, but was: %s", c.expected[i], token)
            }
        }
    }
}
