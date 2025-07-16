package repl

import (
	"strings"
	"fmt"
	"os"
)

type CommandMap struct {
	Name	string
	Description string 
	Callback	func() error
}

// correct this structure 
var CliMap = map[string]CommandMap {
	"exit": {
		Name: "exit",
		Description: "Exit the Pokedex",
		Callback: CommandExit,
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
