package app

import (
	"fmt"
	"os"
	"pokemon/config"
	postgres "pokemon/config/postgre"

	"pokemon/repository"
	"pokemon/services"

	"github.com/gin-gonic/gin"
)

var router = gin.New()

func StartApplication() {
	pokemonRepo := repository.NewPokeRepo(postgres.PSQL.DB.DB)
	app := services.NewPokeUsecase(pokemonRepo)
	config.RegisterApi(router, app)

	port := os.Getenv("APP_PORT")
	router.Run(fmt.Sprintf(":%s", port))
}
