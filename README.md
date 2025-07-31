# 🐾 pokedex-cli

A command-line Pokedex tool written in Go that interacts with the PokeAPI to fetch and explore Pokémon data. It handles API requests, unmarshals JSON responses, supports paginated listings, and displays Pokémon details in an interactive terminal experience. Built using Go’s standard libraries, it includes basic testing and planned caching features for better performance. 

To run the project, clone the repo using `git clone https://github.com/ssd-81/pokedex-cli.git`, then `cd` into the directory and run `go build -o pokedex && ./pokedex`. You can use commands like `next`, `prev`, `pokemon <name>`, and `exit` in the CLI.  

Example:

```bash
$ pokedex
> help
Available commands: next, prev, pokemon <name>, exit

> pokemon pikachu
Name: Pikachu
Height: 4
Weight: 60
Types: electric
...
