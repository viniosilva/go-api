package main

import (
	"github.com/viniosilva/go-api/api"
	"github.com/viniosilva/go-api/app"
	"github.com/viniosilva/go-api/internal/config"
	"github.com/viniosilva/go-api/repository"
)

func main() {
	config, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	healthController := api.NewHealthController()

	catRepository := repository.NewCatRepository()
	catService := app.NewCatApp(catRepository)
	catController := api.NewCatController(catService)

	server := api.NewServer(healthController, catController)

	server.Start(config.Api.Host)
}
