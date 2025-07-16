package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// NewScanner()
	scanner := bufio.NewScanner(os.Stdin)
	// --- for continued blocking ---
	var cleanText string
	for scanner.Scan() {
		cleanText = strings.ToLower(scanner.Text())
		slice := strings.Split(cleanText, " ")
		// fmt.Printf("Your command was: %s", slice[0])
		fmt.Print("Your command was: ", slice[0])
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error while reading input")
	}

}
