package main

import (
	"fmt"
	"os"

	"github.com/utkarshjagtap/pokedexcli/internal/pokeapi"
	"github.com/utkarshjagtap/pokedexcli/internal/pokecache"
)
var commands = map[string]clicommands{}

type clicommands struct{
  name string
  description string
  callback func(*config,string)
}

type config struct{
  next *string
  previous *string
  storage *pokecache.Cache
  pokemon string
  pokedex map[string]pokeapi.PokeDetails
}

func init(){
commands["exit"] = clicommands{
    name: "exit",
    description: "Exit the Pokedex",
    callback: commandExit,
}

commands["help"]=clicommands{
    name: "help",
    description: "Displays a help message",
    callback: commandHelp,
}

commands["map"]=clicommands{
    name: "map",
    description: "Gives you a list of 20 locations to explore",
    callback: commandMap,
}

commands["mapb"]=clicommands{
    name: "mapb",
    description:"Map Back, Give you previous list of 20 locations",
    callback: commandMapBack,
  }


commands["explore"]=clicommands{
    name: "explore <area_name>",
    description: "Explores the given area",
    callback: commandExplore,
}

commands["catch"]=clicommands{
    name: "catch <pokemon>",
    description: "catch the given pokemon",
    callback: commandCatch,
}
commands["inspect"]=clicommands{
    name: "inspect <pokemon>",
    description: "inspect the given pokemon if it is in your pokedex",
    callback: commandinspect,
}
commands["pokedex"]=clicommands{
    name: "pokedex",
    description: "displays list of pokemons that you have caught",
    callback: commandPokedex,
}
}


func commandExit(conf *config, str string) {
  fmt.Printf("Closing the Pokedex... Goodbye!")
  defer os.Exit(0)
}

func commandHelp(conf *config, str string) {
  welcome := `Welcome to the Pokedex!
Usage:

`
  fmt.Println(welcome)
  for _, command := range commands{
    fmt.Println(command.name,": ",command.description)
  }

}
