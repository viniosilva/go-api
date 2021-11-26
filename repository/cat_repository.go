package repository

import "github.com/viniosilva/go-api/model"

//go:generate mockgen -destination=cat_repository_mock.go -package=repository . CatRepository
type CatRepository interface {
	FindCats() []model.Cat
}

type CatRepositoryImpl struct{}

func NewCatRepository() CatRepository {
	return &CatRepositoryImpl{}
}

func (impl CatRepositoryImpl) FindCats() []model.Cat {
	return []model.Cat{}
}
