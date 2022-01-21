package app

import (
	"github.com/viniosilva/go-api/app/cat"
)

type App interface {
	Cat() cat.CatApp
}

type app struct {
	cat cat.CatApp
}

func NewApp(cat cat.CatApp) App {
	return &app{cat: cat}
}

func (app *app) Cat() cat.CatApp { return app.cat }
