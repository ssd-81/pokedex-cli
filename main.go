package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ssd-81/pokedex-cli/internal/pokecache"
	"github.com/ssd-81/pokedex-cli/internal/repl"
)

func main() {
	// NewScanner()
	scanner := bufio.NewScanner(os.Stdin)
	// --- for continued blocking ---
	var cleanText string
	cache := pokecache.NewCache(time.Second * 50)
	c := &repl.Config{"https://pokeapi.co/api/v2/location-area/", "", "", make(map[string]repl.PokemonCatch), cache}
	for {
		fmt.Print("Pokedex > ")

		// doubtful

		commandArgs := repl.CommandArgs{}
		args := []string{}
		commandArgs.Args = args

		//

		if !scanner.Scan() {
			break
		}
		cleanText = strings.ToLower(scanner.Text())
		slice := strings.Split(cleanText, " ")
		command := slice[0]
		// for now , handling for one argument , will extend for multiple args later
		// todo: extend this for handling multiple args
		if len(slice) >= 2 {
			// fmt.Println("command args contains: ---", commandArgs.Args)
			commandArgs.Args = append(commandArgs.Args, slice[1:]...)
			// c.Argument = slice[1]
		}

		if cmd, exists := repl.CliMap[command]; exists {
			if err := cmd.Callback(c, commandArgs); err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error while reading input")
	}

}
