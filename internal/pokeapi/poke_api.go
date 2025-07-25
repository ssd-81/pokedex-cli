package pokeapi

// import ("encoding/json"
// 		"fmt"
// 		"net/http"
// 		"io"
// 		"github.com/ssd-81/pokedex-cli/internal/repl")


// func ListLocations(c *repl.Config) error {
// 	// baseUrl := "https://pokeapi.co/api/v2/location-area/"
// 	resp, err := http.Get(c.Next)
// 	if err != nil {
// 		return fmt.Errorf("Error: ", err)
// 	}
// 	defer resp.Body.Close()
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return fmt.Errorf("Error: ", err)
// 	}
// 	var locations repl.ResultLocations
// 	err = json.Unmarshal(body, &locations)
// 	if err != nil {
// 		return fmt.Errorf("Error:", err)
// 	}

// 	// this is the part where config is required 
// 	c.Next = locations.Next
// 	c.Previous = locations.Previous
// 	for _, loc := range locations.Results {
// 		fmt.Println(loc.Name)
// 	}
// 	return nil

// }

// func ListPrevLocations(c *repl.Config) error {
// 	var prevCall string
// 	if c.Previous != "" {
// 		prevCall = c.Previous
// 	} else {
// 		fmt.Println("you're on the first page")
// 		return nil 
// 	}
// 	resp, err := http.Get(prevCall)
// 	if err != nil {
// 		return fmt.Errorf("Error: ", err)
// 	}
// 	defer resp.Body.Close()
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return fmt.Errorf("Error: ", err)
// 	}
// 	var locations repl.ResultLocations
// 	err = json.Unmarshal(body, &locations)
// 	if err != nil {
// 		return fmt.Errorf("Error:", err)
// 	}
// 	c.Next = locations.Next
// 	c.Previous = locations.Previous
// 	for _, loc := range locations.Results {
// 		fmt.Println(loc.Name)
// 	}
// 	return nil
// }
