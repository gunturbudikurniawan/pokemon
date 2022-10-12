package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"pokemon/models"
	"pokemon/repository/query"
)

type PokemonRepo interface {
	GetAllPokemons() (res models.AllPokemon, err error)
	GetPokemonByName(name string) (res models.GetPokemon, err error)
	GetBattle(start_time, end_time string) (res []models.BattleResponse, err error)
	GetPlayer(BattleID int) (res []models.DetailPlayers, err error)
	GetPokemonScore() (res []models.DetailPlayers, err error)
	PostPokemonData(input models.Pokemon) error
	PostBattlePokemon(input models.BattleInput) (Id int64, err error)
}

func (p *PokeRepo) GetAllPokemons() (res models.AllPokemon, err error) {
	response, err := http.Get("http://pokeapi.co/api/v2/pokemon?limit=100000&offset=0")
	if err != nil {
		return res, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return res, err

	}

	err = json.Unmarshal(responseData, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (p *PokeRepo) GetPokemonByName(name string) (res models.GetPokemon, err error) {
	response, err := http.Get("http://pokeapi.co/api/v2/pokemon/" + name)
	if err != nil {
		return res, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return res, err

	}

	err = json.Unmarshal(responseData, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (p *PokeRepo) PostBattlePokemon(input models.BattleInput) (Id int64, err error) {
	err = p.db.QueryRow(
		query.PostBattle,
		input.Winner,
		input.StartTime,
		input.EndTime,
	).Err()

	if err != nil {
		return Id, err
	}
	return Id, nil
}

func (p *PokeRepo) PostPokemonData(input models.Pokemon) error {
	_, err := p.db.Exec(
		query.PostPokemon,
		input.Name,
		input.BattleID,
		input.Scores,
	)

	if err != nil {
		return err
	}
	return nil
}

func (p *PokeRepo) GetBattle(start_time, end_time string) (res []models.BattleResponse, err error) {
	var qry = query.GetBattle

	if start_time != "" && end_time != "" {
		qry += fmt.Sprintf(" WHERE start_time::date Between '%v' AND '%v' ", start_time, end_time)

	}

	row, err := p.db.Query(
		qry,
	)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		temp := models.BattleResponse{}
		err = row.Scan(
			&temp.BattleID,
			&temp.Winner,
		)
		if err != nil {
			return nil, err
		}

		res = append(res, temp)
	}

	return res, nil
}

func (p *PokeRepo) GetPlayer(BattleID int) (res []models.DetailPlayers, err error) {
	row, err := p.db.Query(
		query.GetPlayers,
		BattleID,
	)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		temp := models.DetailPlayers{}
		err = row.Scan(
			&temp.Name,
			&temp.Scores,
		)
		if err != nil {
			return nil, err
		}

		res = append(res, temp)
	}
	return res, nil
}

func (p *PokeRepo) GetPokemonScore() (res []models.DetailPlayers, err error) {
	row, err := p.db.Query(
		query.GetPokemonScore,
	)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		temp := models.DetailPlayers{}
		err = row.Scan(
			&temp.Name,
			&temp.Scores,
		)
		if err != nil {
			return nil, err
		}
		res = append(res, temp)
	}

	return res, nil
}
