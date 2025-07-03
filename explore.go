package main

import (
	"encoding/json"
	"fmt"

	"github.com/utkarshjagtap/pokedexcli/internal/pokeapi"
)

func commandExplore(conf *config, str string) {
	url := "https://pokeapi.co/api/v2/location-area/" + str

	bytedata, ok := conf.storage.Get(str)
	if !ok {
    var err error
		bytedata, err = pokeapi.SendRequest("explore", url)
		if err != nil {
			fmt.Println("There was an error while exploring this Area", err)
			return
		}

		conf.storage.Add(str, bytedata)
	}

	var data pokeapi.Area

	err := json.Unmarshal(bytedata, &data)
	if err != nil {
		fmt.Println("There was an error unmarshalling", err)
		return
	}

	for _, area := range data.ListOfPokemons {
		fmt.Println(area.Pokemon.Name)
	}

}
