package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
  "github.com/utkarshjagtap/pokedexcli/map"

)

var cute = map[string]cliCommand{}

func main(){

cute["exit"] = cliCommand{
    name: "exit",
    description: "Exit the Pokedex",
    callback: commandExit,
}

cute["help"]=cliCommand{
    name: "help",
    description: "Displays a help message",
    callback: commandHelp,
}

cute["map"]=cliCommand{
    name: "map",
    description: "Gives you a list of 20 locations to explore",
    callback: maps.CommandMap,
}

cute["explore"]=cliCommand{
    name: "explore",
    description: "Explores the given city ",
    callback: commandHelp,
}

  scn := bufio.NewScanner(os.Stdin)
  for {
    fmt.Print("Pokedex >")
    if scn.Scan(){
      clean, err := cleanInput(scn.Text())
      if err != nil{
        fmt.Println(err)
        continue
      }
      cute[clean].callback()
      continue
    }
  }
}

func commandMap() {
}

func cleanInput(text string) (string, error){
  output := strings.ToLower(text)
  switch output {
  case "help":
  case "exit":
  case "map":
  return output, nil

  default:
  return "", fmt.Errorf("Unknown Command %v", output)
    
  }
  return output, nil
}

func commandExit() {
  fmt.Printf("Closing the Pokedex... Goodbye!")
  defer os.Exit(0)
}

func commandHelp() {
  welcome := `Welcome to the Pokedex!
Usage:

`
  fmt.Println(welcome)
  for _, command := range cute{
    fmt.Println(command.name,": ",command.description)
  }

}


type cliCommand struct{
  name string
  description string
  callback func()
}


