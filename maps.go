package main

import (
	"encoding/json"
	"fmt"

	"github.com/utkarshjagtap/pokedexcli/internal/pokeapi"
)

func commandMap(conf *config, str string) {

	url := "https://pokeapi.co/api/v2/location-area/"
	if conf.next != nil {
		url = *conf.next
	}

	var bytesdata []byte
	var err error
	var ok bool

	if bytesdata, ok = conf.storage.Get(url); !ok {

		bytesdata, err = pokeapi.SendRequest("locations", url)
		if err != nil {
			fmt.Sprintln("There was an error while getting locations", err)
			return
		}

		conf.storage.Add(url, bytesdata)
	}

	var result pokeapi.LocationList

	err = json.Unmarshal(bytesdata, &result)
	if err != nil {
		fmt.Println("There was an error while getting locations", err)
		return
	}

	conf.next = result.Next
	conf.previous = result.Previous

	for _, city := range result.Results {
		fmt.Println(city.Location)

	}
}

func commandMapBack(conf *config, str string) {
	url := ""
	if conf.previous == nil {
		fmt.Println("there are no previous pages, try 'help' command to see usage")
		return
	}
	if conf.previous != nil {
		url = *conf.previous
	}
	var bytesdata []byte
	var err error
	var ok bool
	if bytesdata, ok = conf.storage.Get(url); !ok {
		bytesdata, err = pokeapi.SendRequest("locations", url)
		if err != nil {
			fmt.Println("There was an error while getting locations ", err)
			return
		}

		conf.storage.Add(url, bytesdata)
	}

	var result pokeapi.LocationList
	err = json.Unmarshal(bytesdata, &result)
	if err != nil {

		fmt.Println("There was an error while getting locations ", err)
		return
	}
	conf.next = result.Next
	conf.previous = result.Previous

	for _, city := range result.Results {

		fmt.Println(city.Location)

	}
}

