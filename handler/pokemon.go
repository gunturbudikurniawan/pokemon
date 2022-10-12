package handler

import (
	"net/http"
	"pokemon/models"
	"pokemon/services"

	"github.com/gin-gonic/gin"
)


type PokemonHttpServer struct {
	app services.PokeUsecaseInterface
}

func NewPokemonHttpServer(app services.PokeUsecaseInterface) PokemonHttpServer {
	return PokemonHttpServer{app: app}
}

func (p *PokemonHttpServer) PostBattle(c *gin.Context) {
	var req models.RequestBattle
	err := c.BindJSON(&req)
	if err != nil {
		return
	}

	res, err := p.app.PostBattle(req.Pokemons)
	if err != nil {
		return
	}

	c.JSON(http.StatusCreated, res)

}

func (p *PokemonHttpServer) GetAllPokemons(c *gin.Context) {
	data, err := p.app.GetAllPokemons()
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, data)
}

func (p *PokemonHttpServer) GetBattle(c *gin.Context) {
	start := c.Query("start_time")
	end := c.Query("end_time")
	data, err := p.app.GetBattle(start, end)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, data)
}

func (p *PokemonHttpServer) GetPokemonScore(c *gin.Context) {
	data, err := p.app.GetPokemonScore()
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, data)
}
