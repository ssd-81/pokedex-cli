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
		if command == "exit" {
			value, ok := repl.CliMap["exit"]
			if ok {
				if err := value.Callback(); err != nil {
					fmt.Errorf("callback function failed")
				}
			} else {
				fmt.Errorf("this command does not exist")
			}
		}else if command == "help" {
			value , ok := repl.CliMap["help"]
			if ok {
				if err := value.Callback(); err != nil {
					fmt.Errorf("callback function failed")
				}
			} else {
				fmt.Errorf("this command does not exist")
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error while reading input")
	}

}
