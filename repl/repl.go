package repl

import (
	"fmt"
	"os"
	"strings"
)

type CommandMap struct {
	Name        string
	Description string
	Callback    func() error
}

// correct this structure
var CliMap = map[string]CommandMap{
	"exit": {
		Name:        "exit",
		Description: "Exit the Pokedex",
		Callback:    CommandExit,
	},
	"help": {
		Name:        "help",
		Description: "Get help",
		Callback:    CommandHelp,
	},
}

func CleanInput(text string) []string {
	wordSlice := strings.Split(text, " ")
	// fmt.Println(wordSlice)
	for v, word := range wordSlice {
		wordSlice[v] = strings.TrimSpace(word)
	}
	return wordSlice
}

func CommandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil

}

func CommandHelp() error {
	fmt.Println(`Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex`)
	return nil
}
