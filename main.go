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
		// fmt.Print("Your command was: ", slice[0])
		if slice[0] == "exit" {
			err := repl.CommandExit()
			if err != nil {
				fmt.Errorf("Error while exiting pokedex %")
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error while reading input")
	}

}
