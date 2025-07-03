package pokeapi

type LocationList struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Location string `json:"name"`
	} `json:"results"`
}

type Area struct {
	ListOfPokemons []PokeList `json:"pokemon_encounters"`
}

type PokeList struct {
	Pokemon Pokemon `json:"pokemon"`
}


type Pokemon struct {
	Name    string `json:"name"`
}

type PokeDetails struct{
  Name string  `json:"name"`
  Height int `json:"height"`
  Weight int `json:"weight"`
  BaseExperience int `json:"base_experience"`
  Stats []struct{
    BaseStat int `json:"base_stat"`
    Stat struct{
      Name string `json:"name"`
    }`json:"stat"`
  } `json:"stats"`

  Types []struct{
    Type struct{
      Name string `json:"name"`
    } `json:"type"`
  } `json:"types"`

}
