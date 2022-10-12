package models

import "time"

type Battle struct {
	BattleID  int       `json:"battle_id"`
	Winner    string    `json:"winner"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

type RequestBattle struct {
	Pokemons int `json:"pokemons"`
}

type BattleInput struct {
	Winner    string    `json:"winner"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

type BattleResponse struct {
	BattleID int             `json:"battle_id"`
	Winner   string          `json:"winner"`
	Player   []DetailPlayers `json:"player"`
}

type DetailPlayers struct {
	Name   string `json:"name"`
	Scores int    `json:"scores"`
}
