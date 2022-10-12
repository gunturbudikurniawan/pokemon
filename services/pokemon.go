package services

import (
	"fmt"
	"math/rand"
	"pokemon/models"
	"sort"
	"time"
)

type PokemonUsecase interface {
	PostBattle(input int) (res models.BattleResponse, err error)
	GetAllPokemons() (res models.AllPokemon, err error)
	GetBattle(start_time, end_time string) (res []models.BattleResponse, err error)
	GetPokemonScore() (res []models.DetailPlayers, err error)
}

func (p *PokeUsecase) PostBattle(input int) (resp models.BattleResponse, err error) {
	var (
		battleInput models.BattleInput
		fight       = make([]models.GetPokemon, 0)
		idx         = make([]int, 0)
		scores      = make(map[string]int)
		now         = time.Now()
	)

	res, err := p.PokeRepository.GetAllPokemons()
	if err != nil {
		return resp, err
	}

	for i := 0; i < 5; i++ {
		rand.Seed(time.Now().UnixNano())
		min := 0
		max := 15
		Id := rand.Intn(max-min+1) + min
		idx = append(idx, Id)
	}

	for _, v := range idx {
		for i, z := range res.Results {
			if v == i {
				data, err := p.PokeRepository.GetPokemonByName(z.Name)
				if err != nil {
					return resp, err
				}
				fight = append(fight, data)
			}
		}
	}

	for _, poke1 := range fight {
		for _, poke2 := range fight {
			var att1, att2 int
			for _, stat := range poke1.Stats {
				if stat.Stat.Name == "attack" {
					att1 = stat.BaseStat
				}

			}

			for _, stat := range poke2.Stats {
				if stat.Stat.Name == "attack" {
					att2 = stat.BaseStat
				}

			}

			if att1 > att2 {
				scores[poke1.Name] += 3
			} else if att2 > att1 {
				scores[poke2.Name] += 3
			} else {
				scores[poke1.Name]++
				scores[poke2.Name]++

			}
		}

	}

	var maxScore int
	for name, score := range scores {

		if score > maxScore {
			maxScore = score
			battleInput.Winner = name
			battleInput.StartTime = now
			battleInput.EndTime = now.Add(time.Minute * time.Duration(score))

		}
	}

	Id, err := p.PokeRepository.PostBattlePokemon(battleInput)
	if err != nil {
		return resp, err
	}

	type kv struct {
		Key   string
		Value int
	}

	var ss []kv
	for k, v := range scores {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	for _, kv := range ss {
		fmt.Printf("%s, %d\n", kv.Key, kv.Value)
	}

	resp = models.BattleResponse{
		BattleID: int(Id),
		Winner:   ss[0].Key,
	}
	for i, v := range ss {
		dataPlayer := models.Pokemon{
			Name:     v.Key,
			BattleID: int(Id),
			Scores:   5 - i,
		}

		err := p.PokeRepository.PostPokemonData(dataPlayer)
		if err != nil {
			return resp, err
		}

	}

	players, err := p.PokeRepository.GetPlayer(int(Id))
	if err != nil {
		return resp, err
	}

	resp.Player = players

	return resp, nil
}

func (p *PokeUsecase) GetAllPokemons() (res models.AllPokemon, err error) {
	res, err = p.PokeRepository.GetAllPokemons()
	if err != nil {
		return models.AllPokemon{}, err
	}
	return res, nil
}

func (p *PokeUsecase) GetBattle(start_time, end_time string) (res []models.BattleResponse, err error) {
	res, err = p.PokeRepository.GetBattle(start_time, end_time)
	if err != nil {
		return nil, err
	}

	for i, v := range res {
		player, err := p.PokeRepository.GetPlayer(v.BattleID)
		if err != nil {
			return nil, err
		}

		res[i].Player = player
	}

	return res, nil
}

func (p *PokeUsecase) GetPokemonScore() (res []models.DetailPlayers, err error) {
	res, err = p.PokeRepository.GetPokemonScore()
	if err != nil {
		return nil, err
	}
	return res, nil
}
