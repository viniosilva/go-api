package app

import (
	"github.com/viniosilva/go-api/model"
	repo "github.com/viniosilva/go-api/repository"
)

//go:generate mockgen -destination=cat_app_mock.go -package=app . CatApp
type CatApp interface {
	FindCats() []model.Cat
}

type CatAppImpl struct {
	repo repo.CatRepository
}

func NewCatApp(catRepository repo.CatRepository) CatApp {
	return &CatAppImpl{repo: catRepository}
}

func (impl CatAppImpl) FindCats() []model.Cat {
	cats := impl.repo.FindCats()
	return cats
}
