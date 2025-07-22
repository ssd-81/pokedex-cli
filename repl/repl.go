package repl

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type Config struct {
	Next		string
	Previous	string 
}

type CommandMap struct {
	Name        string
	Description string
	Callback    func(c *Config) error
}

type ResultLocations struct {
	Next		string `json:"next"`
	Previous	string `json:"previous"`
	Results []Location `json:"results"` 
}


type Location struct {
    Name    string `json:"name"`
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
	"map": {
		Name:        "map",
		Description: "display next 20 location areas in Pokemon world",
		Callback:    MapCommand,
	},
	"mapb": {
		Name:		"mapb",
		Description: "display previous 20 location areas in Pokemon world",
		Callback: 	MapBCommand,
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

func CommandExit(c *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil

}

func CommandHelp(c *Config) error {
	fmt.Println(`Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex`)
	return nil
}

func MapCommand(c *Config) error {
	// baseUrl := "https://pokeapi.co/api/v2/location-area/"
	resp, err := http.Get(c.Next)
	if err != nil {
		return fmt.Errorf("Error: ", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Error: ", err)
	}
	var locations ResultLocations
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return fmt.Errorf("Error:", err)
	}
	c.Next = locations.Next
	c.Previous = locations.Previous
	// fmt.Println(locations) // this is not what we need; we require the locations to be simply displayed
	for _, loc := range locations.Results {
		fmt.Println(loc)
	}
	return nil

}

func MapBCommand(c *Config) error {
	return nil 
}
