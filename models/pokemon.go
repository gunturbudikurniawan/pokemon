package models

type Pokemon struct {
	PokemonID int    `json:"pokemon_id"`
	Name      string `json:"name"`
	BattleID  int    `json:"battle_id"`
	Scores    int    `json:"scores"`
}

type GetPokemon struct {
	Name  string  `json:"name"`
	Stats []Stats `json:"stats"`
}

type Stats struct {
	BaseStat int  `json:"base_stat"`
	Effort   int  `json:"effort"`
	Stat     Stat `json:"stat"`
}

type Stat struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type AllPokemon struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}
