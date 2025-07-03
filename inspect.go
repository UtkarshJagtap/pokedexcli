package main

import(
  "fmt"
)

func commandinspect(c *config, pokemon string){

  data, ok := c.pokedex[pokemon]
  if !ok{
    fmt.Println("you have not caught",pokemon)
    return
  }
  fmt.Println("Name:", data.Name)
  fmt.Println("Height:", data.Height)
  fmt.Println("Weight:", data.Weight)
  fmt.Println("Stats:")
  for _, stat := range data.Stats{
    fmt.Println(" -", stat.Stat.Name+":", stat.BaseStat)
  }
  fmt.Println("Types:")
  for _, typ := range data.Types{
   fmt.Println("  -", typ.Type.Name) 
  } 

}

func commandPokedex(c *config, str string){
  if len(c.pokedex) == 0{
    fmt.Println("Your Pokedex is empty, try catching some using 'catch' command")
    return
  }
  fmt.Println("Your Pokedex")
  for name, _ := range c.pokedex{
    fmt.Println("-  ",name)
  }
}
