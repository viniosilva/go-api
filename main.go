package main

import (
	"github.com/viniosilva/go-api/api"
	catApp "github.com/viniosilva/go-api/app/cat"
	"github.com/viniosilva/go-api/internal/config"
	"github.com/viniosilva/go-api/internal/gorm"
	"github.com/viniosilva/go-api/model"
	"github.com/viniosilva/go-api/repository"
)

func main() {
	config, err := config.GetConfig("config.yml")
	if err != nil {
		panic(err)
	}

	gorm := gorm.NewGorm(gorm.GormParams{
		Host:     config.MySQL.Host,
		Port:     config.MySQL.Port,
		Database: config.MySQL.Database,
		Username: config.MySQL.Username,
		Password: config.MySQL.Password,
	})
	gorm.Migrate(&model.Cat{})
	gormDB := gorm.GetDB()

	healthController := api.NewHealthController()

	catRepository := repository.NewCatRepository(gormDB)
	catService := catApp.NewCatApp(catRepository)
	catController := api.NewCatController(catService)

	server := api.NewServer(healthController, catController)

	server.Start(config.Api.Host)
}
