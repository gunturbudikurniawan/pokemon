package main

import (
	"pokemon/app"
	config "pokemon/config/postgre"

	"github.com/joho/godotenv"
)


func init() {
	godotenv.Load()
	if err := config.InitPostgre(); err != nil {
		panic(err)
	}
}

func main() {
	app.StartApplication()
}
