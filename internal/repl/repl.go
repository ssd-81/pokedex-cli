package repl

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/ssd-81/pokedex-cli/internal/pokecache"
)

type Config struct {
	Next     string
	Previous string
	Argument string
	PokeDex  map[string]PokemonCatch
	Cache    *pokecache.Cache
}

type CommandMap struct {
	Name        string
	Description string
	Callback    func(c *Config, args CommandArgs) error
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

type PokemonCatch struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
}

// struct for storing the arguments for passing to command callbacks
type CommandArgs struct {
	Args []string
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
	"catch": {
		Name:        "catch",
		Description: "catch a pokemon by throwing a pokeball",
		Callback:    CommandCatch,
	},
	"inspect": {
		Name:        "inspect",
		Description: "takes the name of a Pokemon and prints the name, height, weight, stats and type(s) of the Pokemon",
		Callback:    CommandInspect,
	},
	"pokedex": {
		Name:        "pokedex",
		Description: "show the list of all the caught pokemon",
		Callback:    CommandPokedex,
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

func CommandExit(c *Config, args CommandArgs) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil

}

func CommandHelp(c *Config, args CommandArgs) error {
	fmt.Println(`Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
map: Display next 20 location areas in Pokemon world"
mapb: Display previous 20 location areas in Pokemon world
explore: List of all the Pokemon located in an area"
catch: Catch a pokemon by throwing a pokeball
inspect: Takes the name of a Pokemon and prints the name, height, weight, stats and type(s) of the Pokemon
pokedex: Show the list of all the caught pokemon`)
	return nil
}

func printLocations(locations []Location) {
	for _, location := range locations {
		fmt.Println(location.Name)
	}
}

func MapCommand(c *Config, args CommandArgs) error {
	body, ok := c.Cache.Get(c.Next)
	if !ok {
		resp, err := http.Get(c.Next)
		if err != nil {
			return fmt.Errorf("error: %w", err)
		}
		defer resp.Body.Close()
		newBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("error: %w", err)
		}
		c.Cache.Add(c.Next, newBody)
		var locations ResultLocations
		err = json.Unmarshal(newBody, &locations)
		if err != nil {
			return fmt.Errorf("error: %v", err)
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

func MapBCommand(c *Config, args CommandArgs) error {
	if c.Previous == "" {
		return fmt.Errorf("you are on the first page")
	}
	body, ok := c.Cache.Get(c.Previous)
	if !ok {
		resp, err := http.Get(c.Previous)
		if err != nil {
			return fmt.Errorf("error: %w", err)
		}
		defer resp.Body.Close()
		newBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("error: %w", err)
		}
		c.Cache.Add(c.Previous, newBody)
		var locations ResultLocations
		err = json.Unmarshal(newBody, &locations)
		if err != nil {
			return fmt.Errorf("error: %v	", err)
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

func CommandExplore(c *Config, args CommandArgs) error {

	baseUrl := "https://pokeapi.co/api/v2/location-area/"
	if len(args.Args) != 0 {
		baseUrl += args.Args[len(args.Args)-1] + "/"
	}
	fmt.Println(baseUrl)
	cachedBody, ok := c.Cache.Get(baseUrl)
	if !ok {
		resp, err := http.Get(baseUrl)
		if err != nil {
			return fmt.Errorf("error: %w", err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("error: %w", err)
		}
		c.Cache.Add(baseUrl, body)
		// this is probably where we are messing up
		var encounters Encounters
		err = json.Unmarshal(body, &encounters)
		if err != nil {
			return fmt.Errorf("error: %v", err)
		}

		for _, entry := range encounters.PokemonEncounters {
			fmt.Println(entry.Pokemon.Name)
		}

		return nil
	} else {
		var encounters Encounters
		err := json.Unmarshal(cachedBody, &encounters)
		if err != nil {
			return fmt.Errorf("error while retrieving data from cache %w", err)
		}
		for _, entry := range encounters.PokemonEncounters {
			fmt.Println(entry.Pokemon.Name)
		}
	}
	return nil
}

func CommandCatch(c *Config, args CommandArgs) error {
	if len(args.Args) == 0 {
		fmt.Println("usage: catch <pokemon name>")
		return nil
	}
	fmt.Printf("Throwing a Pokeball at %v...\n", args.Args[0])
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", args.Args[0])
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	var pokemonData PokemonCatch
	err = json.Unmarshal(body, &pokemonData)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}
	catchPower := r.Intn(1 + r.Intn(20))
	if catchPower*60+100 > pokemonData.BaseExperience {
		fmt.Printf("%s was caught...\n", pokemonData.Name)
		c.PokeDex[pokemonData.Name] = pokemonData
	} else {
		fmt.Printf("%s escaped!\n", pokemonData.Name)
	}

	return nil

}

func CommandInspect(c *Config, args CommandArgs) error {
	if len(args.Args) == 0 {
		return fmt.Errorf("command usage > inspect <pokemon_name>")
	}
	pokemonName := args.Args[0]
	val, ok := c.PokeDex[pokemonName]
	if ok {
		fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\nBaseExperience: %d\n", val.Name, val.Height, val.Weight, val.BaseExperience)

		return nil
	} else {
		return fmt.Errorf("pokemon %q not found in the pokedex", pokemonName)
	}

}

func CommandPokedex(c *Config, args CommandArgs) error {
	fmt.Println("Your Pokedex:")
	for key, _ := range c.PokeDex {
		fmt.Println(key)
	}
	return nil
}
