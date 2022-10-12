package config

import (
	"pokemon/handler"
	"pokemon/services"

	"github.com/gin-gonic/gin"
)

func RegisterApi(r *gin.Engine, app services.PokeUsecaseInterface) {
	pokeSrv :=  handler.NewPokemonHttpServer(app)
	api := r.Group("/pokemon")
	{
		api.POST("/battle", pokeSrv.PostBattle)
		api.GET("/", pokeSrv.GetAllPokemons)
		api.GET("/battles", pokeSrv.GetBattle)
		api.GET("/scores", pokeSrv.GetPokemonScore)
	}
}
