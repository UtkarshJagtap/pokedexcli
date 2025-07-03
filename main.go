package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/utkarshjagtap/pokedexcli/internal/pokeapi"
	"github.com/utkarshjagtap/pokedexcli/internal/pokecache"
)

func main() {

	configuration := config{
		storage: pokecache.NewCache(time.Second * 5),
    pokedex: map[string]pokeapi.PokeDetails{},
	}

	startREPL(&configuration)
}

func startREPL(c *config) {
	scn := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		if scn.Scan() {
      text := scn.Text()
			clean, err := cleanInput(text)
			if err != nil {
				fmt.Println(err)
				continue
			}
      value, ok := commands[clean[0]]
      if !ok {
        fmt.Println("invalid command '", clean[0], "' try 'help' to see usage")
        continue
      }
      var arg string
      if len(clean) > 1{
      value.callback(c, clean[1])
			continue
      }
      value.callback(c, arg)
      continue
		}
	}
}

func cleanInput(text string) ([]string, error) {
  
	output := strings.Fields(strings.ToLower(text))
  if len(output) > 2{
    return nil, fmt.Errorf("invalid command '%v', use 'help' command to see usage", text)
  }
 	return output, nil
}
