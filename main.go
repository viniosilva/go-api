package main

import (
	"github.com/viniosilva/go-api/api"
	"github.com/viniosilva/go-api/api/cat"
	"github.com/viniosilva/go-api/api/health"
	"github.com/viniosilva/go-api/app"
	catApp "github.com/viniosilva/go-api/app/cat"
	"github.com/viniosilva/go-api/internal/config"
	"github.com/viniosilva/go-api/internal/gorm"
	"github.com/viniosilva/go-api/model"
	"github.com/viniosilva/go-api/repository"
	catRepo "github.com/viniosilva/go-api/repository/cat"
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

	repo := repository.NewRepository(catRepo.NewRepository(gormDB))
	application := app.NewApp(catApp.NewApp(repo.Cat()))
	server := api.NewApi(health.NewApi(), cat.NewApi(application.Cat()))

	server.Start(config.Api.Host)
}
