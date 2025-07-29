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
	Argument string
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

type Encounters struct {
	Name              string    `json:"name"`
	PokemonEncounters []Details `json:"pokemon_encounters"`
}

type Details struct {
	Pokemon Pokemon `json:"pokemon"`
}

type Pokemon struct {
	Name string `json:"name"`
}

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
	"explore": {
		Name:        "explore",
		Description: "list of all the Pokemon located in an area",
		Callback:    CommandExplore,
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
			return fmt.Errorf("Error: %v", err)
		}
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
			return fmt.Errorf("Error: %v	", err)
		}
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
		fmt.Println("from cache!")
		if err != nil {
			return fmt.Errorf("error while retrieving data from cache %w", err)
		}
		printLocations(locations.Results)
		return nil
	}
}

func CommandExplore(c *Config) error {
	// the most important part is extracting the location area from the user command
	// how are we going to do that? think for some time.
	// make request to this : https://pokeapi.co/api/v2/location-area/{id or name}/
	// derive the pokemon from the response
	// print the pokemons
	baseUrl := "https://pokeapi.co/api/v2/location-area/"
	if c.Argument != "" {
		baseUrl += c.Argument
	}
	resp, err := http.Get(baseUrl)
	if err != nil {
		return fmt.Errorf("Error: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Error: %w", err)
	}

	// this is probably where we are messing up
	var encounters Encounters
	err = json.Unmarshal(body, &encounters)
	if err != nil {
		return fmt.Errorf("Error: %v", err)
	}

	for _, entry := range encounters.PokemonEncounters {
		fmt.Println(entry.Pokemon.Name)
	}

	return nil
}
