package main

import (
	"encoding/json"
	"fmt"

	"github.com/utkarshjagtap/pokedexcli/internal/pokeapi"
)

func commandCatch(c *config, pokemon string){
  url := "https://pokeapi.co/api/v2/pokemon/"+pokemon

  fmt.Println("Throwing a Pokeball at",pokemon+"...")

  bytedata, ok := c.storage.Get(pokemon)
  if !ok{
    var err error
    bytedata, err = pokeapi.SendRequest("pokemon",url)
    if err != nil{
      fmt.Println("There was an error while catching", pokemon, err)
      return
    }
    c.storage.Add(pokemon, bytedata)
  }


  var pokedetails pokeapi.PokeDetails
  err := json.Unmarshal(bytedata, &pokedetails)
  if err != nil {
    fmt.Println("There was an error unmarshalling this pokemon", err)
  }

  c.pokedex[pokemon] = pokedetails

  fmt.Println(pokemon,"was caught")






}
