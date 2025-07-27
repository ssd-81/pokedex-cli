package repl

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"github.com/ssd-81/pokedex-cli/internal/pokecache"
)

type Config struct {
	Next     string
	Previous string
	Cache    *pokecache.Cache
}

type CommandMap struct {
	Name        string
	Description string
	Callback    func(c *Config) error
}

type ResultLocations struct {
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []Location `json:"results"`
}

type Location struct {
	Name string `json:"name"`
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
		Name:        "mapb",
		Description: "display previous 20 location areas in Pokemon world",
		Callback:    MapBCommand,
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

func printLocations(locations []Location) {
	for _, location := range locations {
		fmt.Println(location.Name)
	}
}

func MapCommand(c *Config) error {
	body, ok := c.Cache.Get(c.Next)
	if !ok {
		resp, err := http.Get(c.Next)
		if err != nil {
			return fmt.Errorf("Error: %w", err)
		}
		defer resp.Body.Close()
		newBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("Error: %w", err)
		}
		c.Cache.Add(c.Next, newBody)
		var locations ResultLocations
		err = json.Unmarshal(newBody, &locations)
		if err != nil {
			return fmt.Errorf("Error:", err)
		}
		// for testing
		fmt.Println("\n\n")
		fmt.Println(c.Next)
		fmt.Println(c.Previous)
		fmt.Println("\n\n")
		//
		c.Next = locations.Next
		c.Previous = locations.Previous
		for _, loc := range locations.Results {
			fmt.Println(loc.Name)
		}
		return nil
	} else {
		var locations ResultLocations
		err := json.Unmarshal(body, &locations)
		c.Next = locations.Next
		c.Previous = locations.Previous
		fmt.Println("from cache!\n\n\n")
		if err != nil {
			return fmt.Errorf("error while retrieving data from cache %w", err)
		}
		printLocations(locations.Results)
		return nil
	}
}

func MapBCommand(c *Config) error {
	if c.Previous == "" {
		return fmt.Errorf("you are on the first page")
	}
	body, ok := c.Cache.Get(c.Previous)
	if !ok {
		resp, err := http.Get(c.Previous)
		if err != nil {
			return fmt.Errorf("Error: %w", err)
		}
		defer resp.Body.Close()
		newBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("Error: %w", err)
		}
		c.Cache.Add(c.Previous, newBody)
		var locations ResultLocations
		err = json.Unmarshal(newBody, &locations)
		if err != nil {
			return fmt.Errorf("Error:", err)
		}
		// for testing
		fmt.Println("\n\n")
		fmt.Println(c.Next)
		fmt.Println(c.Previous)
		fmt.Println("\n\n")
		//
		c.Next = locations.Next
		c.Previous = locations.Previous
		for _, loc := range locations.Results {
			fmt.Println(loc.Name)
		}
		return nil
	} else {
		var locations ResultLocations
		err := json.Unmarshal(body, &locations)
		c.Next = locations.Next
		c.Previous = locations.Previous
		fmt.Println("from cache!\n\n\n")
		if err != nil {
			return fmt.Errorf("error while retrieving data from cache %w", err)
		}
		printLocations(locations.Results)
		return nil
	}
}
