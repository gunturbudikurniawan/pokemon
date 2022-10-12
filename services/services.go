package services

import "pokemon/repository"

type PokeUsecase struct {
	PokeRepository repository.PokeRepoInterface
}

type PokeUsecaseInterface interface {
	PokemonUsecase
}

func NewPokeUsecase(pokeRepo repository.PokeRepoInterface) PokeUsecaseInterface {
	return &PokeUsecase{
		PokeRepository: pokeRepo,
	}
}
