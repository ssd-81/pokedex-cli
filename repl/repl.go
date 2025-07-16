package repl

import (
	"strings"
	"fmt"
	"os"
)

type CommandMap struct {
	name	string
	description string 
	callback	func() error
}

// correct this structure 
var cliMap = map[string]CommandMap {
	"exit": {
		name: "exit",
		description: "Exit the Pokedex",
		callback: CommandExit,
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
