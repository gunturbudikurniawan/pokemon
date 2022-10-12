package repository

import "database/sql"

type PokeRepo struct {
	db *sql.DB
}

type PokeRepoInterface interface {
	PokemonRepo
}

func NewPokeRepo(db *sql.DB) *PokeRepo {
	return &PokeRepo{db: db}
}
