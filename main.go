package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"github.com/ssd-81/pokedex-cli/internal/repl"
	"github.com/ssd-81/pokedex-cli/internal/pokecache"
)

func main() {
	// NewScanner()
	scanner := bufio.NewScanner(os.Stdin)
	// --- for continued blocking ---
	var cleanText string
	c := &repl.Config{"https://pokeapi.co/api/v2/location-area/", ""}
	cache := pokecache.NewCache(time.Second * 5)
	for {

		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			break
		}
		cleanText = strings.ToLower(scanner.Text())
		slice := strings.Split(cleanText, " ")
		command := slice[0]

		if cmd, exists := repl.CliMap[command]; exists {
			if err := cmd.Callback(c); err != nil {
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
