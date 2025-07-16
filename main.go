package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ssd-81/pokedex-cli/repl"
)

func main() {
	// NewScanner()
	scanner := bufio.NewScanner(os.Stdin)
	// --- for continued blocking ---
	var cleanText string
	for scanner.Scan() {
		cleanText = strings.ToLower(scanner.Text())
		slice := strings.Split(cleanText, " ")
		command := slice[0]
		// fmt.Print("Your command was: ", slice[0])
		if cmd, exists := repl.CliMap[command]; exists {
			if err := cmd.Callback(); err != nil {
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
